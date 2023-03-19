package storage

import (
	apiStorageService "wbHTTPServer/storage-service/api/user"

	"gorm.io/gorm"
)

type UserStorageService struct {
	apiStorageService.UserStorageServiceServer

	db *gorm.DB
}

func NewUserStorageService(db *gorm.DB) *UserStorageService {
	return &UserStorageService{db: db}
}
