package entity

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrGenerateUUID = status.Error(codes.Internal, "failed generate uuid")
)