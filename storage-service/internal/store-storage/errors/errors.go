package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrCreateStore = status.Error(codes.Internal, "create store error")
	ErrGetStore    = status.Error(codes.Internal, "get store error")
	ErrDeleteStore = status.Error(codes.Internal, "delete store error")
	ErrUpdateStore = status.Error(codes.Internal, "store update error")
)
