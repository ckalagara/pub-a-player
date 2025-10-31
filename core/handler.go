package core

import (
	"context"
	"net/http"

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
	//TODO implement me
	panic("implement me")
}

func (h handlerImpl) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h handlerImpl) UploadAttachment(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h handlerImpl) Shutdown(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h handlerImpl) Health(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
