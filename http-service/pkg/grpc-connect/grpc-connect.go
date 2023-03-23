package grpc

import (
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetServiceConnection(logger zerolog.Logger, serviceHost string, opts ...grpc.DialOption) *grpc.ClientConn {
	logger.Debug().Msgf("getting service connection for host: %s", serviceHost)

	conn, err := grpc.Dial(
		serviceHost,
		append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))...,
	)
	if nil != err {
		logger.Panic().Err(err).Msg("")
	}

	logger.Info().Msgf("connection state for host '%s': %s", serviceHost, conn.GetState())

	return conn
}
