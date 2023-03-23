package storage

import (
	"context"
	apiCore "wbHTTPServer/storage-service/api/core"
	apiStorageService "wbHTTPServer/storage-service/api/store"
	"wbHTTPServer/storage-service/internal/db/entity"
	errService "wbHTTPServer/storage-service/internal/store-storage/errors"
	"wbHTTPServer/storage-service/pkg/converters"

	"github.com/rs/zerolog"
)

func (receiver *StoreStorageService) UpdateStore(ctx context.Context, request *apiStorageService.UpdateStoreRequest) (*apiCore.EmptyMessage, error) {
	logger := zerolog.Ctx(ctx)

	storeUID, err := converters.String2UUID(logger, &request.Uid)
	if err != nil {
		logger.Error().Err(err).Msg("failed convert string to uuid")

		return nil, err
	}

	storeProfileEntity := &entity.StoreProfileEntity{
		UID:     storeUID,
		Name:    request.Name,
		Address: request.Address,
		Working: request.Working,
		Owner:   request.Owner,
	}
	if dbErr := receiver.db.Model(&storeProfileEntity).Updates(storeProfileEntity).Error; dbErr != nil {
		logger.Error().Err(dbErr).Msg("failed update store")

		return nil, errService.ErrUpdateStore
	}

	logger.Trace().Msgf("store successfully updated. storeUID: %s", request.GetUid())

	return &apiCore.EmptyMessage{}, nil
}
