package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Migrate(db *gorm.DB) *gorm.DB {
	return db.AutoMigrate(
		&User{},
		&Podcast{},
		&Episode{},
	)
}
