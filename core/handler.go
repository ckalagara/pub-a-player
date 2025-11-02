package core

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/mail"

	"github.com/ckalagara/pub-a-player/commons"
	"gorm.io/gorm"
)

type Handler interface {
	GetPlayer(http.ResponseWriter, *http.Request)
	UpdatePlayer(http.ResponseWriter, *http.Request)
	UploadAttachment(http.ResponseWriter, *http.Request)
	Shutdown(ctx context.Context) error
	Health(ctx context.Context) error
}

func NewHandler(ctx context.Context, db *gorm.DB) Handler {

	return &handlerImpl{
		s: newStorePostgres(ctx, db),
	}
}

type handlerImpl struct {
	s store
}

func (h handlerImpl) GetPlayer(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("X-Pub-Email")
	_, err := mail.ParseAddress(email)
	if err != nil {
		writeErrResponse(w, http.StatusBadRequest, err)
		return
	}

	var p Player
	p, err = h.s.Get(r.Context(), "email", email)
	if err != nil {
		writeErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	writeJson(w, http.StatusOK, p)

}

func (h handlerImpl) UpdatePlayer(w http.ResponseWriter, r *http.Request) {

	var p Player
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeErrResponse(w, http.StatusBadRequest, err)
	}

	if err := h.s.Update(r.Context(), p); err != nil {
		writeErrResponse(w, http.StatusInternalServerError, err)
	}
}

func (h handlerImpl) UploadAttachment(_ http.ResponseWriter, _ *http.Request) {
	//TODO implement me
	return
}

func (h handlerImpl) Shutdown(ctx context.Context) error {
	return h.s.Shutdown(ctx)
}

func (h handlerImpl) Health(ctx context.Context) error {
	return h.s.Health(ctx)
}

func writeErrResponse(w http.ResponseWriter, code int, err error) {
	writeJson(w, code, commons.GenericResponse{
		Status:      "Error",
		Description: err.Error(),
	})
}

func writeJson(w http.ResponseWriter, code int, d any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(d); err != nil {
		log.Printf("failed to write response: %v", err)
		return
	}
}
