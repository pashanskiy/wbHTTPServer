package interceptor

import (
	"context"

	"github.com/pereslava/grpc_zerolog"
	"github.com/rs/zerolog"

	"google.golang.org/grpc"
)

const ReceiveRequest = "Receive request: %++v"

func GrpcServerWithInterceptors(logger zerolog.Logger) []grpc.ServerOption {
	unaryInterceptors := grpc.ChainUnaryInterceptor(
		grpc_zerolog.NewUnaryServerInterceptor(logger),
		grpcLogInterceptor(logger),
	)
	streamInterceptors := grpc.ChainStreamInterceptor(
		grpc_zerolog.NewStreamServerInterceptor(logger),
	)
	return []grpc.ServerOption{unaryInterceptors, streamInterceptors}
}

func grpcLogInterceptor(logger zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		newCtx := logger.WithContext(ctx)

		logger.Debug().Msgf(ReceiveRequest, req)

		return handler(newCtx, req)
	}
}
