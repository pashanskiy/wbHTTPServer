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
	handleCreateRoute = httpRouter.Route{
		Pattern: servicePath + "/create",
		Method:  "POST",
	}
)

func (s *UserHTTPService) handleCreate(w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context())

	request := models.HandleCreateJSON{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error().Err(err).Msg("failed unmarshall request")

		httpErrors.ErrStatusBadRequest(w, err)
		return
	}

	_, err := s.userStorageApi.CreateUser(r.Context(), &proto.CreateUsersRequest{
		Surname:    request.Surname,
		Name:       request.Name,
		Secondname: request.Secondname,
		Age:        request.Age,
	})

	if err != nil {
		logger.Error().Err(err).Msg("failed create user request")

		httpErrors.ErrStatusInternal(w, err)
		return
	}

	logger.Info().Msgf("user successfully created")

	w.WriteHeader(http.StatusOK)
}
