package app

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/cyansnbrst/gesture-guru/sso-service/config"
)

// App struct
type App struct {
	config *config.Config
	logger *zap.Logger
	db     *pgxpool.Pool
}

// New app constructor
func NewServer(cfg *config.Config, logger *zap.Logger, db *pgxpool.Pool) *App {
	return &App{
		config: cfg,
		logger: logger,
		db:     db,
	}
}

// Run starts the app
func (a *App) Run() error {
	addr := fmt.Sprintf(":%d", a.config.App.GRPC.Port)
	server := grpc.NewServer()

	a.RegisterServices(server)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	a.logger.Info("starting gRPC server",
		zap.String("addr", addr),
		zap.String("env", a.config.App.Env),
	)

	shutDownError := make(chan error)

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
			shutDownError <- nil
		case <-time.After(a.config.App.GRPC.Timeout):
			server.Stop()
			shutDownError <- fmt.Errorf("forced shutdown after timeout")
		}
	}()

	if err := server.Serve(listener); err != nil {
		return err
	}

	if err := <-shutDownError; err != nil {
		return err
	}

	a.logger.Info("stopped gRPC server",
		zap.String("addr", addr),
	)

	return nil
}
