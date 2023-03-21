package storage

import (
	"context"
	apiStorageService "wbHTTPServer/storage-service/api/store"
	"wbHTTPServer/storage-service/internal/db/entity"
	errService "wbHTTPServer/storage-service/internal/store-storage/errors"
	"wbHTTPServer/storage-service/pkg/converters"

	"github.com/rs/zerolog"
)

func (receiver *StoreStorageService) GetStores(ctx context.Context, request *apiStorageService.StoreInfo) (*apiStorageService.GetStoresResponse, error) {
	logger := zerolog.Ctx(ctx)

	storeUID, err := converters.String2UUID(logger, request.Uid)
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

	storeProfileEntities := &[]entity.StoreProfileEntity{}

	if dbErr := receiver.db.Find(storeProfileEntities, storeProfileEntity).Error; dbErr != nil {
		logger.Error().Err(dbErr).Msg("failed get store")

		return nil, errService.ErrGetStore
	}

	storesList := make([]*apiStorageService.StoreInfo, len(*storeProfileEntities))

	for idx, store := range *storeProfileEntities {
		stringUID := store.UID.String()
		storesList[idx] = &apiStorageService.StoreInfo{
			Uid:     &stringUID,
			Name:    store.Name,
			Address: store.Address,
			Working: store.Working,
			Owner:   store.Owner,
		}
	}

	logger.Trace().Msgf("store successfully was taken. stores len: %d", len(storesList))

	return &apiStorageService.GetStoresResponse{
		StoresInfo: storesList,
	}, nil
}
