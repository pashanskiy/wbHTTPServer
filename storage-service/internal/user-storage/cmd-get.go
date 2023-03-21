package storage

import (
	"context"
	apiStorageService "wbHTTPServer/storage-service/api/user"
	"wbHTTPServer/storage-service/internal/db/entity"
	errService "wbHTTPServer/storage-service/internal/user-storage/errors"

	"github.com/rs/zerolog"
)

func (receiver *UserStorageService) GetUsers(ctx context.Context, request *apiStorageService.UserInfo) (*apiStorageService.GetUsersResponse, error) {
	logger := zerolog.Ctx(ctx)

	userUID, err := receiver.string2UUID(logger, request.Uid)
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

	userProfileEntities := &[]entity.UserProfileEntity{}

	if dbErr := receiver.db.Find(userProfileEntities, userProfileEntity).Error; dbErr != nil {
		logger.Error().Err(dbErr).Msg("failed get user")

		return nil, errService.ErrGetUser
	}

	usersList := make([]*apiStorageService.UserInfo, len(*userProfileEntities))

	for idx, user := range *userProfileEntities {
		stringUID := user.UID.String()
		date := user.CreatedAt.UnixNano()
		usersList[idx] = &apiStorageService.UserInfo{
			Uid:                   &stringUID,
			Surname:               user.Surname,
			Name:                  user.Name,
			Secondname:            user.Secondname,
			Age:                   user.Age,
			RegisterDateTimestamp: &date,
		}
	}

	logger.Trace().Msgf("user successfully was taken. users len: %d", len(usersList))

	return &apiStorageService.GetUsersResponse{
		UsersInfo: usersList,
	}, nil
}
