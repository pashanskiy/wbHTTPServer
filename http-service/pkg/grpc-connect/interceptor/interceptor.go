package interceptor

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/pereslava/grpc_zerolog"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	grpcMetadata "google.golang.org/grpc/metadata"
)

const HeaderUserInfo = "authorization"

func GetGRPCInterceptors(logger zerolog.Logger) []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithUnaryInterceptor(NewClientInterceptor().UnaryHeaderForwarder()),
		grpc.WithStreamInterceptor(NewClientInterceptor().UnaryStreamHeaderForwarder()),

		grpc.WithChainUnaryInterceptor(grpc_zerolog.NewUnaryClientInterceptor(logger)),
		grpc.WithChainStreamInterceptor(grpc_zerolog.NewStreamClientInterceptor(logger)),
	}
}

type ClientInfoInterceptor struct{}

func NewClientInterceptor() *ClientInfoInterceptor {
	return &ClientInfoInterceptor{}
}

func (receiver *ClientInfoInterceptor) UnaryHeaderForwarder() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		return invoker(receiver.interceptor(ctx), method, req, reply, cc, opts...)
	}
}

func (receiver *ClientInfoInterceptor) UnaryStreamHeaderForwarder() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		return streamer(receiver.interceptor(ctx), desc, cc, method, opts...)
	}
}

func (receiver *ClientInfoInterceptor) interceptor(ctx context.Context) context.Context {
	logger := zerolog.Ctx(ctx)

	headerValue := metadata.ExtractIncoming(ctx).Get(HeaderUserInfo)
	logger.Trace().Msgf("headerValue: %s", headerValue)

	if headerValue == "" {
		logger.Trace().Msg("user info was not found")

		return ctx
	}

	logger.Trace().Msg("user info token found - forward it further in the request...")

	md := grpcMetadata.New(map[string]string{HeaderUserInfo: headerValue})
	return grpcMetadata.NewOutgoingContext(ctx, md)
}
