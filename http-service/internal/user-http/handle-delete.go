package http

import (
	"encoding/json"
	"net/http"

	httpErrors "wbHTTPServer/http-service/internal/user-http/errors"
	models "wbHTTPServer/http-service/internal/user-http/models"
	httpRouter "wbHTTPServer/http-service/pkg/http-router"
	proto "wbHTTPServer/storage-service/api/user"

	"github.com/rs/zerolog"
)

var (
	handleDeleteRoute = httpRouter.Route{
		Pattern: servicePath + "/delete",
		Method:  http.MethodDelete,
	}
)

func (s *UserHTTPService) handleDelete(w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context())

	request := models.DeleteUserJSON{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error().Err(err).Msg("failed unmarshall request")

		httpErrors.ErrStatusBadRequest(w, err)
		return
	}

	_, err := s.userStorageApi.DeleteUser(r.Context(), &proto.DeleteUserRequest{
		Uid: request.UID,
	})

	if err != nil {
		logger.Error().Err(err).Msg("failed deleted user request")

		httpErrors.ErrStatusInternal(w, err)
		return
	}

	logger.Info().Msgf("user successfully deleted")

	w.WriteHeader(http.StatusOK)
}
