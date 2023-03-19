package storage

import (
	"context"
	apiCore "wbHTTPServer/storage-service/api/core"
	apiStorageService "wbHTTPServer/storage-service/api/user"

	"github.com/rs/zerolog"
)

func (receiver *UserStorageService) CreateUser(ctx context.Context, request *apiStorageService.UserInfo) (*apiCore.EmptyMessage, error) {
	logger := zerolog.Ctx(ctx)

	logger.Info().Msgf("create request: %++v", request)

	return &apiCore.EmptyMessage{}, nil
}
