package core

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	UUID      string  `gorm:"type:uuid;default:gen_random_uuid();uniqueIndex"`
	Name      string  `gorm:"type:varchar(100);not null"`
	Email     string  `gorm:"type:varchar(100);uniqueIndex;not null"`
	Age       int     `gorm:"not null"`
	Team      string  `gorm:"type:varchar(50)"`
	Score     float64 `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Upload struct {
	ID         uint   `gorm:"primaryKey"`
	Email      string `gorm:"not null;index:idx_email_type_filename,unique"` // part of unique composite key
	UploadType string `gorm:"not null;index:idx_email_type_filename,unique"` // part of unique composite key
	Filename   string `gorm:"not null;index:idx_email_type_filename,unique"` // part of unique composite key
	Data       []byte `gorm:"not null"`
}
