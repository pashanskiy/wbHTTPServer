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
	handleUpdateRoute = httpRouter.Route{
		Pattern: servicePath + "/update",
		Method:  http.MethodPatch,
	}
)

func (s *StoreHTTPService) handleUpdate(w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context())

	request := models.HandleUpdateJSON{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error().Err(err).Msg("failed unmarshall request")

		httpErrors.ErrStatusBadRequest(w, err)
		return
	}

	_, err := s.storeStorageApi.UpdateStore(r.Context(), &proto.UpdateStoreRequest{
		Uid:     request.UID,
		Name:    request.Name,
		Address: request.Address,
		Working: request.Working,
		Owner:   request.Owner,
	})

	if err != nil {
		logger.Error().Err(err).Msg("failed update store request")

		httpErrors.ErrStatusInternal(w, err)
		return
	}

	logger.Info().Msgf("store successfully updated")

	w.WriteHeader(http.StatusOK)
}
