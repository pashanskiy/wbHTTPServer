package interceptor

import (
	"net/http"

	"github.com/rs/zerolog"
)

type Interceptor struct {
	handler http.Handler

	logger zerolog.Logger
}

func (i *Interceptor) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//add logger to request context
	logger := i.addUserInfoToLogger(r)
	r = r.WithContext(logger.WithContext(r.Context()))

	i.handler.ServeHTTP(w, r)

	i.logger.Info().Msgf("Request done. method: %s, path: %s, from: %s", r.Method, r.URL, r.RemoteAddr)
}

func AddToMux(logger zerolog.Logger, handler http.Handler) *Interceptor {
	return &Interceptor{
		logger:  logger,
		handler: handler,
	}
}

func (i *Interceptor)addUserInfoToLogger(r *http.Request) zerolog.Logger {

	//MARK: так то тут должно выдергивать инфу о юзере из токена, но для примера и так сойдет
	loggerContext := i.logger.With().Str("userUID", r.Header.Get("user-uid"))

	return loggerContext.Logger()
}
