package storage

import (
	"context"
	apiStorageService "wbHTTPServer/storage-service/api/user"

	"github.com/rs/zerolog"
)

func (receiver *UserStorageService) GetUser(ctx context.Context, request *apiStorageService.GetUserInfo) (*apiStorageService.UserInfo, error) {
	logger := zerolog.Ctx(ctx)

	logger.Info().Msgf("get request: %++v", request)

	return &apiStorageService.UserInfo{}, nil
}
