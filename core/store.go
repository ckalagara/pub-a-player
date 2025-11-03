package core

import (
	"context"
	"fmt"
	"log"

	"github.com/ckalagara/pub-a-player/commons"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type store interface {
	Get(ctx context.Context, field, value string) (Player, error)
	Update(ctx context.Context, p Player) error
	Shutdown(ctx context.Context) error
	Health(ctx context.Context) error
	Upload(ctx context.Context, upload Upload) error
}

func newStorePostgres(_ context.Context, db *gorm.DB) store {
	err := db.AutoMigrate(&Player{}, &Upload{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &storePostgresImpl{
		client: db,
	}
}

type storePostgresImpl struct {
	client *gorm.DB
}

func (s storePostgresImpl) Upload(ctx context.Context, upload Upload) error {
	return s.client.WithContext(ctx).Create(&upload).Error
}

func (s storePostgresImpl) Shutdown(ctx context.Context) error {
	uClient, err := s.client.WithContext(ctx).DB()
	if err != nil {
		return err
	}
	return uClient.Close()
}

func (s storePostgresImpl) Health(ctx context.Context) error {
	uClient, err := s.client.WithContext(ctx).DB()
	if err != nil {
		return err
	}
	return uClient.Ping()
}

func (s storePostgresImpl) Get(ctx context.Context, field, value string) (Player, error) {
	var p Player
	err := s.client.WithContext(ctx).Where(fmt.Sprintf(qryByField, field), value).First(&p).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Player{}, errors.Wrap(commons.ErrPlayerNotFound, err.Error())
		}
		return Player{}, err
	}
	return p, nil
}

func (s storePostgresImpl) Update(ctx context.Context, p Player) error {
	r := s.client.WithContext(ctx).Model(&Player{}).Where(fmt.Sprintf(qryByField, "email"), p.Email).Updates(p)

	if r.Error != nil {
		return r.Error
	}

	if r.RowsAffected == 0 {
		err := s.client.WithContext(ctx).Create(&p).Error
		return errors.Wrap(err, fmt.Sprintf("failed to create record for email %s", p.Email))
	}

	return nil
}

const (
	qryByField = "%s = ?"
)
