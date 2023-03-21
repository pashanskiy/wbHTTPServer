package storage

import (
	apiStorageService "wbHTTPServer/storage-service/api/store"

	"gorm.io/gorm"
)

type StoreStorageService struct {
	apiStorageService.StoreStorageServiceServer

	db *gorm.DB
}

func NewStoreStorageService(db *gorm.DB) *StoreStorageService {
	return &StoreStorageService{db: db}
}
