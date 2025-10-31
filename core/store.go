package core

import (
	"context"

	"gorm.io/gorm"
)

type store interface {
	Get(ctx context.Context, field, value string) (string, error)
	Update(ctx context.Context, p Player) error
	Shutdown(ctx context.Context) error
	Health(ctx context.Context) error
}

func newStorePostgres(ctx context.Context, db *gorm.DB) store {
	return &storePostgresImpl{
		client: db,
	}
}

type storePostgresImpl struct {
	client *gorm.DB
}

func (s storePostgresImpl) Shutdown(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s storePostgresImpl) Health(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s storePostgresImpl) Get(ctx context.Context, field, value string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s storePostgresImpl) Update(ctx context.Context, p Player) error {
	//TODO implement me
	panic("implement me")
}
