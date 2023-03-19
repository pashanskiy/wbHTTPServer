package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserProfileEntity struct {
	UID string `gorm:"column:uid; primaryKey"`

	Surname    string `gorm:"column:surname"`
	Name       string `gorm:"column:name"`
	Secondname string `gorm:"column:secondname"`
	Age        uint32 `gorm:"column:age"`

	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (UserProfileEntity) TableName() string { return "users_profiles" }
