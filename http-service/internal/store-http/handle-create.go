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
	handleCreateRoute = httpRouter.Route{
		Pattern: servicePath + "/create",
		Method:  http.MethodPost,
	}
)

func (s *StoreHTTPService) handleCreate(w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context())

	request := models.HandleCreateJSON{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error().Err(err).Msg("failed unmarshall request")

		httpErrors.ErrStatusBadRequest(w, err)
		return
	}

	_, err := s.storeStorageApi.CreateStore(r.Context(), &proto.CreateStoreRequest{
		Name:    request.Name,
		Address: request.Address,
		Working: request.Working,
		Owner:   request.Owner,
	})

	if err != nil {
		logger.Error().Err(err).Msg("failed create store request")

		httpErrors.ErrStatusInternal(w, err)
		return
	}

	logger.Info().Msgf("store successfully created")

	w.WriteHeader(http.StatusOK)
}
