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
	handleDeleteRoute = httpRouter.Route{
		Pattern: servicePath + "/delete",
		Method:  http.MethodDelete,
	}
)

func (s *StoreHTTPService) handleDelete(w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context())

	request := models.DeleteStoreJSON{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error().Err(err).Msg("failed unmarshall request")

		httpErrors.ErrStatusBadRequest(w, err)
		return
	}

	_, err := s.storeStorageApi.DeleteStore(r.Context(), &proto.DeleteStoreRequest{
		Uid: request.UID,
	})

	if err != nil {
		logger.Error().Err(err).Msg("failed deleted store request")

		httpErrors.ErrStatusInternal(w, err)
		return
	}

	logger.Info().Msgf("store successfully deleted")

	w.WriteHeader(http.StatusOK)
}
