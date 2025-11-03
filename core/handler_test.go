package core

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ckalagara/pub-a-player/commons"
	"github.com/stretchr/testify/mock"
)

var (
	testPlayer = Player{
		Name:  "John Doe",
		Email: "johndoe@example.com",
		Age:   29,
		Team:  "Warriors",
		Score: 1500.75,
	}
)

func Test_handlerImpl_GetPlayer(t *testing.T) {
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		preTest  func(t *testing.T, m *mockstore)
		PostTest func(t *testing.T, m *mockstore, w *httptest.ResponseRecorder)
	}{
		{
			name: "get player",
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					r := httptest.NewRequest("GET", "/player", nil)
					r.Header.Set("x-pub-email", testPlayer.Email)
					return r
				}(),
			},
			preTest: func(t *testing.T, m *mockstore) {
				// Get(ctx context.Context, field, value string) (Player, error)
				m.On("Get", mock.IsType(context.Background()), "email", testPlayer.Email).Return(testPlayer, nil)
			},
			PostTest: func(t *testing.T, m *mockstore, w *httptest.ResponseRecorder) {
				m.AssertExpectations(t)
				if w.Code != 200 {
					t.Errorf("Expected 200, got %d", w.Code)
					return
				}
				b := w.Body.Bytes()
				var p Player

				err := json.Unmarshal(b, &p)
				if err != nil {
					t.Error(err)
					return
				}

				if p.Email != testPlayer.Email {
					t.Errorf("Expected %s, got %s", testPlayer.Email, p.Email)
					return
				}
			},
		},
		{
			name: "invalid email hdr",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/player", nil),
			},
			preTest: func(t *testing.T, m *mockstore) {},
			PostTest: func(t *testing.T, m *mockstore, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusBadRequest {
					t.Errorf("Expected 400, got %d", w.Code)
				}
			},
		},
		{
			name: "missing player info",
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					r := httptest.NewRequest("GET", "/player", nil)
					r.Header.Set("x-pub-email", testPlayer.Email)
					return r
				}(),
			},
			preTest: func(t *testing.T, m *mockstore) {
				// Get(ctx context.Context, field, value string) (Player, error)
				m.On("Get", mock.IsType(context.Background()), "email", testPlayer.Email).Return(Player{}, commons.ErrPlayerNotFound)
			},
			PostTest: func(t *testing.T, m *mockstore, w *httptest.ResponseRecorder) {
				m.AssertExpectations(t)
				if w.Code != 500 {
					t.Errorf("Expected 500, got %d", w.Code)
					return
				}
				b := w.Body.Bytes()
				var p commons.GenericResponse

				err := json.Unmarshal(b, &p)
				if err != nil {
					t.Error(err)
					return
				}

				if p.Description != commons.PlayerNotFound {
					t.Errorf("Expected %s, got %s", commons.PlayerNotFound, p.Description)
					return
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := new(mockstore)
			tt.preTest(t, m)
			h := handlerImpl{
				s: m,
			}
			h.GetPlayer(tt.args.w, tt.args.r)

			tt.PostTest(t, m, tt.args.w)
		})
	}
}
