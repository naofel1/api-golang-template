package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/naofel1/api-golang-template/internal/configs"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// Server provides HTTP routing and handler dependencies.
type Server struct {
	config *configs.Config
	logger *otelzap.Logger

	// server listens for incoming HTTP requests and routes them to the correct
	// handler.
	server *http.Server

	// errorStream carries errors from the running server across process
	// boundaries. This is particularly helpful for monitoring errors from
	// multiple concurrent sources, e.g. server errors and OS interrupts.
	errorStream chan error
}

// NewServer returns a new hexagonal server configured using the provided Config.
func NewServer(
	logger *otelzap.Logger,
	handler *gin.Engine,
	cfg *configs.Config,
) *Server {
	server := Server{
		config: cfg,
		logger: logger,
		server: &http.Server{
			Addr:         fmt.Sprintf(":%v", cfg.Host.Port),
			Handler:      handler,
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
		},
		errorStream: make(chan error, 1),
	}

	return &server
}

// Run starts the Server in a new goroutine, forwarding any errors to its
// errorStream. Run blocks until either a server error or an OS interrupt
// occurs. In the case of an interrupt, Run first attempts to shut down
// gracefully.
func (s *Server) Run(ctx context.Context) error {
	s.logger.Ctx(ctx).Info("Starting server",
		zap.String("Server Address", s.config.Host.Address),
		zap.Int("Server Port", s.config.Host.Port),
	)

	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.errorStream <- fmt.Errorf("ListenAndServe: %w", err)
		}
	}()

	select {
	case err := <-s.errorStream:
		return fmt.Errorf("server error: %w", err)
	case sig := <-newSignalHandler():
		s.logger.Ctx(ctx).Info("Received signal",
			zap.Any("Signal:", sig),
		)

		if err := s.GracefulShutdown(); err != nil {
			return fmt.Errorf("graceful shutdown: %w", err)
		}

		return nil
	}
}

// GracefulShutdown closes the Server with the grace period specified in its
// config.
func (s *Server) GracefulShutdown() error {
	s.logger.Info("Shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), s.config.Server.ShutdownGracePeriod)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	return nil
}

// ServeHTTP calls the ServeHTTP method of the Server's underlying handler,
// passing through its ResponseWriter and Request. This is router-agnostic and
// makes the server exceptionally easy to test.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.server.Handler.ServeHTTP(w, r)
}

func newSignalHandler() <-chan os.Signal {
	signalStream := make(chan os.Signal, 1)
	signal.Notify(signalStream, os.Interrupt, syscall.SIGTERM)

	return signalStream
}
