package storage

import (
	"context"
	apiCore "wbHTTPServer/storage-service/api/core"
	apiStorageService "wbHTTPServer/storage-service/api/user"
	"wbHTTPServer/storage-service/internal/db/entity"
	errService "wbHTTPServer/storage-service/internal/user-storage/errors"

	"github.com/rs/zerolog"
)

func (receiver *UserStorageService) CreateUser(ctx context.Context, request *apiStorageService.CreateUsersRequest) (*apiCore.EmptyMessage, error) {
	logger := zerolog.Ctx(ctx)

	userProfileEntity := &entity.UserProfileEntity{
		Surname:    &request.Surname,
		Name:       request.Name,
		Secondname: request.Secondname,
		Age:        request.Age,
	}

	if err := receiver.db.Create(userProfileEntity).Error; err != nil {
		logger.Error().Err(err).Msg("failed create user")

		return nil, errService.ErrCreateUser
	}

	logger.Trace().Msgf("user successfully created. uid: %s", userProfileEntity.UID.String())

	return &apiCore.EmptyMessage{}, nil
}
