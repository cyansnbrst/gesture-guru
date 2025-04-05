package app

import (
	"google.golang.org/grpc"

	grpcapp "github.com/cyansnbrst/gesture-guru/sso-service/internal/auth/delivery/grpc"
	"github.com/cyansnbrst/gesture-guru/sso-service/internal/auth/repository"
	"github.com/cyansnbrst/gesture-guru/sso-service/internal/auth/usecase"
)

func (a *App) RegisterServices(server *grpc.Server) {
	authRepo := repository.NewAuthRepo(a.db)
	authUC := usecase.NewAuthUseCase(a.logger, a.config, authRepo)
	grpcapp.NewAuthServer(server, authUC)
}
