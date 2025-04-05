package logger

import "go.uber.org/zap"

func LogError(logger *zap.Logger, op string, err error, message string) {
	logger.With(zap.String("op", op)).Error(message, zap.Error(err))
}
