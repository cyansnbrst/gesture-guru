package app

import (
	"google.golang.org/grpc"

	grpcapp "github.com/cyansnbrst/gesture-guru/gestures-service/internal/gestures/delivery/grpc"
	"github.com/cyansnbrst/gesture-guru/gestures-service/internal/gestures/repository"
	"github.com/cyansnbrst/gesture-guru/gestures-service/internal/gestures/usecase"
)

func (a *App) RegisterServices(server *grpc.Server) {
	gesturesRepo := repository.NewGesturesRepo(a.db)
	gesturesUC := usecase.NewGesturesUseCase(a.logger, gesturesRepo)
	grpcapp.NewGesturesServer(server, gesturesUC)
}
