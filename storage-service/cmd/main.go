package main

import (
	"fmt"
	"net"

	apiStoreStorageService "wbHTTPServer/storage-service/api/store"
	apiUserStorageService "wbHTTPServer/storage-service/api/user"
	config "wbHTTPServer/storage-service/config"
	dbService "wbHTTPServer/storage-service/internal/db"
	envParse "wbHTTPServer/storage-service/internal/env-parse"
	grpcInterceptors "wbHTTPServer/storage-service/internal/interceptor"
	loggerService "wbHTTPServer/storage-service/internal/logger"
	storeStorageService "wbHTTPServer/storage-service/internal/store-storage"
	userStorageService "wbHTTPServer/storage-service/internal/user-storage"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// defaults for -ldflags
var ServiceName = "Storage"
var BuildDate = "nil"

func main() {
	var serviceConfig config.Config
	envParse.ReadEnvConfig(log.Logger, &serviceConfig)

	logger := loggerService.InitRootLogger(
		loggerService.ParseEnvLoggerEnv(serviceConfig.LogRootLevel),
		ServiceName,
		BuildDate,
	)

	db := dbService.NewDB(logger).Gorm

	grpcServer := grpc.NewServer(grpcInterceptors.GrpcServerWithInterceptors(logger)...)

	apiUserStorageService.RegisterUserStorageServiceServer(grpcServer, userStorageService.NewUserStorageService(db))
	apiStoreStorageService.RegisterStoreStorageServiceServer(grpcServer, storeStorageService.NewStoreStorageService(db))

	runService(serviceConfig, grpcServer)
}

func runService(serviceConfig config.Config, grpcServer *grpc.Server) {
	if serviceConfig.GrpcReflectionEnabled {
		reflection.Register(grpcServer)
	}

	listen, err := net.Listen(
		"tcp4",
		fmt.Sprintf("%s:%d", serviceConfig.GrpcServerHost, serviceConfig.GrpcServerPort),
	)
	if err != nil {
		log.Panic().Err(err).Msg("")
	}

	log.Info().Msgf("gRPC listen: %s", listen.Addr().String())

	if err = grpcServer.Serve(listen); err != nil {
		log.Panic().Err(err).Msg("")
	}
}
