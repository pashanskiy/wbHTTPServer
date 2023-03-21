package entity

import (
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/rs/zerolog"

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

func (entity *StoreProfileEntity) BeforeCreate(tx *gorm.DB) (err error) {
	logger := zerolog.Ctx(tx.Statement.Context)

	uuid, err := uuid.NewV4()
	if err != nil {
		logger.Error().Err(err).Msg("failed generate uuid")

		return ErrGenerateUUID
	}

	entity.UID = &uuid
	return
}
