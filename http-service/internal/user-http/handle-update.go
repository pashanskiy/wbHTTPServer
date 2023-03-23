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
	handleUpdateRoute = httpRouter.Route{
		Pattern: servicePath + "/update",
		Method:  http.MethodPatch,
	}
)

func (s *UserHTTPService) handleUpdate(w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context())

	request := models.HandleUpdateJSON{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error().Err(err).Msg("failed unmarshall request")

		httpErrors.ErrStatusBadRequest(w, err)
		return
	}

	_, err := s.userStorageApi.UpdateUser(r.Context(), &proto.UpdateUsersRequest{
		Uid:        request.UID,
		Surname:    request.Surname,
		Name:       request.Name,
		Secondname: request.Secondname,
		Age:        request.Age,
	})

	if err != nil {
		logger.Error().Err(err).Msg("failed update user request")

		httpErrors.ErrStatusInternal(w, err)
		return
	}

	logger.Info().Msgf("user successfully updated")

	w.WriteHeader(http.StatusOK)
}
