package storage

import (
	apiStorageService "wbHTTPServer/storage-service/api/user"

)

type UserStorageService struct {
	apiStorageService.UserStorageServiceServer
}

func NewUserStorageService() *UserStorageService {
	return &UserStorageService{}
}
