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

func (receiver *StoreStorageService) DeleteStore(ctx context.Context, request *apiStorageService.DeleteStoreRequest) (*apiCore.EmptyMessage, error) {
	logger := zerolog.Ctx(ctx)

	storeUID, err := converters.String2UUID(logger, &request.Uid)
	if err != nil {
		logger.Error().Err(err).Msg("failed convert string to uuid")

		return nil, err
	}

	storeProfileEntity := &entity.StoreProfileEntity{
		UID: storeUID,
	}

	if dbErr := receiver.db.Delete(storeProfileEntity, storeProfileEntity).Error; dbErr != nil {
		logger.Error().Err(dbErr).Msg("failed delete store")

		return nil, errService.ErrDeleteStore
	}

	logger.Trace().Msgf("store successfully deleted. storeUID: %s", storeUID.String())

	return &apiCore.EmptyMessage{}, nil
}
