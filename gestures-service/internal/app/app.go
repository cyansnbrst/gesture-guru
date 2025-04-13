package app

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cyansnbrst/gesture-guru/gestures-service/config"
	"github.com/cyansnbrst/gesture-guru/gestures-service/internal/interceptors"
)

// App struct
type App struct {
	config *config.Config
	logger *zap.Logger
	db     *pgxpool.Pool
}

// New app constructor
func NewApp(cfg *config.Config, logger *zap.Logger, db *pgxpool.Pool) *App {
	return &App{
		config: cfg,
		logger: logger,
		db:     db,
	}
}

// Run starts the app
func (a *App) Run() error {
	server := a.newGRPCServer()
	a.RegisterServices(server)

	listener, err := a.createListener()
	if err != nil {
		return err
	}

	a.logger.Info("starting gRPC server",
		zap.String("addr", listener.Addr().String()),
		zap.String("env", a.config.App.Env),
	)

	shutdownErr := a.handleGracefulShutdown(server)

	if err := server.Serve(listener); err != nil {
		return err
	}

	if err := <-shutdownErr; err != nil {
		return err
	}

	a.logger.Info("stopped gRPC server",
		zap.String("addr", listener.Addr().String()),
	)

	return nil
}

// New GRPC server with interceptors
func (a *App) newGRPCServer() *grpc.Server {
	recoveryOpts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p any) (err error) {
			a.logger.Error("panic recovered", zap.Any("panic", p))
			return status.Errorf(codes.Internal, "internal server error")
		}),
	}

	return grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.UnaryAuthInterceptor(a.config.App.JWTSecretKey),
			grpc_recovery.UnaryServerInterceptor(recoveryOpts...),
		),
	)
}

// Creates a listener
func (a *App) createListener() (net.Listener, error) {
	addr := fmt.Sprintf(":%d", a.config.App.GRPC.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on %s: %w", addr, err)
	}
	return listener, nil
}

// Handles a server graceful shutdown
func (a *App) handleGracefulShutdown(server *grpc.Server) <-chan error {
	shutdownErr := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		sig := <-quit

		a.logger.Info("shutting down server...",
			zap.String("signal", sig.String()),
		)

		stopped := make(chan struct{})
		go func() {
			server.GracefulStop()
			close(stopped)
		}()

		select {
		case <-stopped:
			shutdownErr <- nil
		case <-time.After(a.config.App.GRPC.Timeout):
			server.Stop()
			shutdownErr <- fmt.Errorf("forced shutdown after timeout")
		}
	}()

	return shutdownErr
}
