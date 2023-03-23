package http

import (
	"encoding/json"
	"net/http"

	httpErrors "wbHTTPServer/http-service/internal/store-http/errors"
	models "wbHTTPServer/http-service/internal/store-http/models"
	httpRouter "wbHTTPServer/http-service/pkg/http-router"
	proto "wbHTTPServer/storage-service/api/store"

	"github.com/rs/zerolog"
)

var (
	handleGetRoute = httpRouter.Route{
		Pattern: servicePath + "/get",
		Method:  http.MethodPost,
	}
)

func (s *StoreHTTPService) handleGet(w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context())

	request := models.StoreInfoJSON{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error().Err(err).Msg("failed unmarshall request")

		httpErrors.ErrStatusBadRequest(w, err)
		return
	}

	grpcResponse, err := s.storeStorageApi.GetStores(r.Context(), &proto.StoreInfo{
		Uid:     request.UID,
		Name:    request.Name,
		Address: request.Address,
		Working: request.Working,
		Owner:   request.Owner,
	})

	if err != nil {
		logger.Error().Err(err).Msg("failed get stores request")

		httpErrors.ErrStatusInternal(w, err)
		return
	}

	storesInfo := make([]models.StoreInfoJSON, len(grpcResponse.GetStoresInfo()))

	for idx, storeInfo := range grpcResponse.GetStoresInfo() {
		storesInfo[idx] = models.StoreInfoJSON{
			UID:     storeInfo.Uid,
			Name:    storeInfo.Name,
			Address: storeInfo.Address,
			Working: storeInfo.Working,
			Owner:   storeInfo.Owner,
		}

	}

	logger.Info().Msgf("store was getted successfully")

	if err = json.NewEncoder(w).Encode(&models.StoreInfoListJSON{StoresInfo: storesInfo}); err != nil {
		httpErrors.ErrStatusInternal(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
