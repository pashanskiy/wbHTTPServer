package HttpRouter

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
)

type HttpRouter struct {
	mux *http.ServeMux

	logger zerolog.Logger

	routesMap map[Route]func(http.ResponseWriter, *http.Request)
}

type Route struct {
	Pattern string
	Method  string
}

type RouteHandler struct {
	Route       Route
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
}

func (hr *HttpRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handlerFunc, ok := hr.routesMap[Route{Pattern: r.URL.Path, Method: r.Method}]
	if !ok {
		errMessage := fmt.Sprintf("route %s %s is not found", r.Method, r.URL.Path)
		hr.logger.Error().Msgf(errMessage)

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errMessage))

		return
	}

	handlerFunc(w, r)

}

func New(logger zerolog.Logger, mux *http.ServeMux) *HttpRouter {
	return &HttpRouter{
		logger:    logger,
		mux:       mux,
		routesMap: make(map[Route]func(http.ResponseWriter, *http.Request)),
	}
}

func (hr *HttpRouter) SetRoutes(routes []RouteHandler) {
	for _, r := range routes {

		_, ok := hr.routesMap[r.Route]
		if ok {
			hr.logger.Fatal().Msgf("duplicate http route %s %s", r.Route.Method, r.Route.Pattern)
		}

		hr.routesMap[r.Route] = r.HandlerFunc
	}

	hr.logger.Info().Msgf("routes setted: %d", len(routes))
}
