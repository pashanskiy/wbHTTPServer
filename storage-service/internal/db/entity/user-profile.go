package entity

import (
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/rs/zerolog"

	"gorm.io/gorm"
)

type UserProfileEntity struct {
	UID *uuid.UUID `gorm:"column:uid; type:uuid; primaryKey"`

	Surname    *string `gorm:"column:surname"`
	Name       *string `gorm:"column:name"`
	Secondname *string `gorm:"column:secondname"`
	Age        *int32  `gorm:"column:age"`

	CreatedAt time.Time      `gorm:"column:created_at; <-:create"`
	UpdatedAt time.Time      `gorm:"column:updated_at; autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (UserProfileEntity) TableName() string { return "users_profiles" }

func (entity *UserProfileEntity) BeforeCreate(tx *gorm.DB) (err error) {
	logger := zerolog.Ctx(tx.Statement.Context)

	uuid, err := uuid.NewV4()
	if err != nil {
		logger.Error().Err(err).Msg("failed generate uuid")

		return ErrGenerateUUID
	}

	entity.UID = &uuid
	return
}
