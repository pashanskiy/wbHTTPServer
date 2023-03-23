package http

import (
	"io"
	"net/http"

	httpRouter "wbHTTPServer/http-service/pkg/http-router"

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

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error().Err(err).Msg("failed read request body")
	}

	logger.Info().Msgf("user successfully created")

	w.WriteHeader(http.StatusOK)

	w.Write([]byte(reqBody))
}
