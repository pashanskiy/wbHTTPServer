package storage

import (
	"context"
	apiCore "wbHTTPServer/storage-service/api/core"
	apiStorageService "wbHTTPServer/storage-service/api/user"
	errService "wbHTTPServer/storage-service/internal/user-storage/errors"

	"wbHTTPServer/storage-service/internal/db/entity"

	"github.com/gofrs/uuid/v5"
	"github.com/rs/zerolog"
)

func (receiver *UserStorageService) CreateUser(ctx context.Context, request *apiStorageService.CreateUsersRequest) (*apiCore.EmptyMessage, error) {
	logger := zerolog.Ctx(ctx)

	userUID, err := uuid.NewV4()
	if err != nil {
		logger.Error().Err(err).Msg("failed generate uuid")

		return nil, errService.ErrCreateUser
	}

	userProfileEntity := &entity.UserProfileEntity{
		UID:        &userUID,
		Surname:    &request.Surname,
		Name:       request.Name,
		Secondname: request.Secondname,
		Age:        request.Age,
	}

	if err = receiver.db.Create(userProfileEntity).Error; err != nil {
		logger.Error().Err(err).Msg("failed create user")

		return nil, errService.ErrCreateUser
	}

	logger.Trace().Msgf("user successfully created. uid: %s", userUID.String())

	return &apiCore.EmptyMessage{}, nil
}
