package storage

import (
	"context"
	apiCore "wbHTTPServer/storage-service/api/core"
	apiStorageService "wbHTTPServer/storage-service/api/user"
	"wbHTTPServer/storage-service/internal/db/entity"
	errService "wbHTTPServer/storage-service/internal/user-storage/errors"

	"github.com/rs/zerolog"
)

func (receiver *UserStorageService) UpdateUser(ctx context.Context, request *apiStorageService.UpdateUsersRequest) (*apiCore.EmptyMessage, error) {
	logger := zerolog.Ctx(ctx)

	userUID, err := receiver.string2UUID(logger, &request.Uid)
	if err != nil {
		logger.Error().Err(err).Msg("failed convert string to uuid")

		return nil, err
	}

	userProfileEntity := &entity.UserProfileEntity{
		UID:        userUID,
		Surname:    request.Surname,
		Name:       request.Name,
		Secondname: request.Secondname,
		Age:        request.Age,
	}

	if dbErr := receiver.db.Save(userProfileEntity).Error; dbErr != nil {
		logger.Error().Err(dbErr).Msg("failed update user")

		return nil, errService.ErrUpdateUser
	}

	logger.Trace().Msgf("user successfully updated. userUID: %s", request.GetUid())

	return &apiCore.EmptyMessage{}, nil
}
