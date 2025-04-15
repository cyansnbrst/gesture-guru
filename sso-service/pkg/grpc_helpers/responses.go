package grpchelpers

import (
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleError(logger *zap.Logger, err error, msg string, code codes.Code) error {
	logger.Error(msg, zap.Error(err))
	return status.Error(code, msg)
}
