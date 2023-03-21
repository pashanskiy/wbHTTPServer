package entity

import (
	"time"

	"github.com/gofrs/uuid/v5"

	"gorm.io/gorm"
)

type StoreProfileEntity struct {
	UID *uuid.UUID `gorm:"column:uid; type:uuid; primaryKey"`

	Name    *string `gorm:"column:name"`
	Address *string `gorm:"column:address"`
	Working *bool   `gorm:"column:working"`
	Owner   *string `gorm:"column:owner"`

	CreatedAt time.Time      `gorm:"column:created_at; <-:create"`
	UpdatedAt time.Time      `gorm:"column:updated_at; autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (StoreProfileEntity) TableName() string { return "stores_profiles" }
