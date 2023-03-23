package interceptor

import (
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/rs/zerolog"
)

var HeadersToParse = []string{"authorization"}

type Interceptor struct {
	handler http.Handler

	logger zerolog.Logger
}

func (i *Interceptor) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//add logger to request context
	logger := i.addUserInfoToLogger(r)
	r = r.WithContext(logger.WithContext(r.Context()))

	r = i.addUserInfoFromHeaderToGRPCContext(r)
	i.handler.ServeHTTP(w, r)

	i.logger.Info().Msgf("Request done. method: %s, path: %s, from: %s", r.Method, r.URL, r.RemoteAddr)
}

func AddToMux(logger zerolog.Logger, handler http.Handler) *Interceptor {
	return &Interceptor{
		logger:  logger,
		handler: handler,
	}
}

func (i *Interceptor) addUserInfoToLogger(r *http.Request) zerolog.Logger {

	//MARK: так то тут должно выдергивать инфу о юзере из токена, но для примера и так сойдет
	loggerContext := i.logger.With().Str("authorization", r.Header.Get("authorization"))

	return loggerContext.Logger()
}

func (i *Interceptor) addUserInfoFromHeaderToGRPCContext(r *http.Request) *http.Request {
	ctx := r.Context()
	for _, key := range HeadersToParse {
		value := r.Header.Get(key)
		if value == "" {
			i.logger.Warn().Msgf("empty header: %s", key)
			continue
		}
		ctx = metadata.ExtractIncoming(ctx).Set(key, value).ToIncoming(ctx)
	}

	return r.WithContext(ctx)
}
