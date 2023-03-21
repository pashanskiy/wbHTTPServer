package storage

import (
	"context"
	apiCore "wbHTTPServer/storage-service/api/core"
	apiStorageService "wbHTTPServer/storage-service/api/user"
	"wbHTTPServer/storage-service/internal/db/entity"
	errService "wbHTTPServer/storage-service/internal/user-storage/errors"

	"github.com/rs/zerolog"
)

func (receiver *UserStorageService) DeleteUser(ctx context.Context, request *apiStorageService.DeleteUserRequest) (*apiCore.EmptyMessage, error) {
	logger := zerolog.Ctx(ctx)

	userUID, err := receiver.string2UUID(logger, &request.Uid)
	if err != nil {
		logger.Error().Err(err).Msg("failed convert string to uuid")

		return nil, err
	}

	userProfileEntity := &entity.UserProfileEntity{
		UID: userUID,
	}

	if dbErr := receiver.db.Delete(userProfileEntity, userProfileEntity).Error; dbErr != nil {
		logger.Error().Err(dbErr).Msg("failed delete user")

		return nil, errService.ErrDeleteUser
	}

	logger.Trace().Msgf("user successfully deleted. userUID: %s", userUID.String())

	return &apiCore.EmptyMessage{}, nil
}
