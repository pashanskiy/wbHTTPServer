package storage

import (
	"context"
	apiCore "wbHTTPServer/storage-service/api/core"
	apiStorageService "wbHTTPServer/storage-service/api/store"
	errService "wbHTTPServer/storage-service/internal/store-storage/errors"

	"wbHTTPServer/storage-service/internal/db/entity"

	"github.com/rs/zerolog"
)

func (receiver *StoreStorageService) CreateStore(ctx context.Context, request *apiStorageService.CreateStoreRequest) (*apiCore.EmptyMessage, error) {
	logger := zerolog.Ctx(ctx)

	storeProfileEntity := &entity.StoreProfileEntity{
		Name:    &request.Name,
		Address: request.Address,
		Working: request.Working,
		Owner:   request.Owner,
	}

	if err := receiver.db.Create(storeProfileEntity).Error; err != nil {
		logger.Error().Err(err).Msg("failed create store")

		return nil, errService.ErrCreateStore
	}

	logger.Trace().Msgf("store successfully created. uid: %s", storeProfileEntity.UID.String())

	return &apiCore.EmptyMessage{}, nil
}
