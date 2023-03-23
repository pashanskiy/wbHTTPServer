package main

import (
	"fmt"
	"net"
	"net/http"
	config "wbHTTPServer/http-service/config"
	envParse "wbHTTPServer/http-service/internal/env-parse"
	loggerService "wbHTTPServer/http-service/internal/logger"
	userHTTPService "wbHTTPServer/http-service/internal/user-http"
	httpRouter "wbHTTPServer/http-service/pkg/http-router"
	httpInterceptors "wbHTTPServer/http-service/pkg/http-router/interceptor"

	"github.com/rs/zerolog/log"
)

// defaults for -ldflags
var ServiceName = "HTTP"
var BuildDate = "nil"

func main() {
	var serviceConfig config.Config
	envParse.ReadEnvConfig(log.Logger, &serviceConfig)

	logger := loggerService.InitRootLogger(
		loggerService.ParseEnvLoggerEnv(serviceConfig.LogRootLevel),
		ServiceName,
		BuildDate,
	)

	mux := httpRouter.New(logger, http.NewServeMux())
	userHTTPService.NewUserHTTPService(mux).SetRoutes()

	runService(serviceConfig, httpInterceptors.AddToMux(logger, mux))
}

func runService(serviceConfig config.Config, handler http.Handler) {
	listen, err := net.Listen(
		"tcp4",
		fmt.Sprintf("%s:%d", serviceConfig.HTTPServerHost, serviceConfig.HTTPServerPort),
	)
	if err != nil {
		log.Panic().Err(err).Msg("")
	}

	log.Info().Msgf("HTTP listen: %s", listen.Addr().String())

	if err = http.Serve(listen, handler); err != nil {
		log.Panic().Err(err).Msg("")
	}
}
