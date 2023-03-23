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
	handleGetRoute = httpRouter.Route{
		Pattern: servicePath + "/get",
		Method:  http.MethodPost,
	}
)

func (s *UserHTTPService) handleGet(w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context())

	request := models.UserInfoJSON{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error().Err(err).Msg("failed unmarshall request")

		httpErrors.ErrStatusBadRequest(w, err)
		return
	}

	grpcResponse, err := s.userStorageApi.GetUsers(r.Context(), &proto.UserInfo{
		Uid:                   request.UID,
		Surname:               request.Surname,
		Name:                  request.Name,
		Secondname:            request.Secondname,
		Age:                   request.Age,
		RegisterDateTimestamp: request.RegisterDateTimestamp,
	})

	if err != nil {
		logger.Error().Err(err).Msg("failed get users request")

		httpErrors.ErrStatusInternal(w, err)
		return
	}

	usersInfo := make([]models.UserInfoJSON, len(grpcResponse.GetUsersInfo()))

	for idx, userInfo := range grpcResponse.GetUsersInfo() {
		usersInfo[idx] = models.UserInfoJSON{
			UID:                    userInfo.Uid,
			Surname:                userInfo.Surname,
			Name:                   userInfo.Name,
			Secondname:             userInfo.Secondname,
			Age:                    userInfo.Age,
			RegisterDateTimestamp: userInfo.RegisterDateTimestamp,
		}

	}

	logger.Info().Msgf("user was getted successfully")

	if err = json.NewEncoder(w).Encode(&models.UserInfoListJSON{UsersInfo: usersInfo}); err != nil {
		httpErrors.ErrStatusInternal(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
