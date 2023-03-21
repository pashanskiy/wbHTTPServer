package storage

import (
	apiStorageService "wbHTTPServer/storage-service/api/user"
	errService "wbHTTPServer/storage-service/internal/user-storage/errors"

	"github.com/gofrs/uuid/v5"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type UserStorageService struct {
	apiStorageService.UserStorageServiceServer

	db *gorm.DB
}

func NewUserStorageService(db *gorm.DB) *UserStorageService {
	return &UserStorageService{db: db}
}

func (service UserStorageService) string2UUID(logger *zerolog.Logger, stringUUID *string) (*uuid.UUID, error) {
	if stringUUID == nil {

		return nil, nil
	}

	rawuuid, err := uuid.FromString(*stringUUID)
	if err != nil {
		logger.Error().Err(err).Msg("failed parse userUID")

		return nil, errService.ErrInvalidUID
	}

	return &rawuuid, nil
}
