// Package api provides the HTTP server and API handler for Hetty.
package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"go.uber.org/zap"
)

const (
	defaultReadTimeout  = 10 * time.Second
	defaultWriteTimeout = 10 * time.Second
	defaultIdleTimeout  = 60 * time.Second
)

// Server represents the Hetty HTTP API server.
type Server struct {
	httpServer *http.Server
	logger     *zap.SugaredLogger
	addr       string
}

// Config holds configuration options for the API server.
type Config struct {
	// Addr is the TCP address for the server to listen on, in the form "host:port".
	Addr string
	// Logger is the structured logger used by the server.
	Logger *zap.SugaredLogger
}

// NewServer creates and configures a new API server instance.
func NewServer(cfg Config) (*Server, error) {
	if cfg.Addr == "" {
		cfg.Addr = "127.0.0.1:8080"
	}

	if cfg.Logger == nil {
		return nil, fmt.Errorf("api: logger is required")
	}

	s := &Server{
		addr:   cfg.Addr,
		logger: cfg.Logger,
	}

	mux := http.NewServeMux()
	s.registerRoutes(mux)

	s.httpServer = &http.Server{
		Addr:         cfg.Addr,
		Handler:      mux,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
		IdleTimeout:  defaultIdleTimeout,
	}

	return s, nil
}

// registerRoutes sets up the HTTP routes for the API server.
func (s *Server) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/health", s.handleHealth)
}

// handleHealth responds with a simple health check payload.
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}

// Start begins listening and serving HTTP requests.
func (s *Server) Start() error {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("api: failed to listen on %s: %w", s.addr, err)
	}

	s.logger.Infow("API server listening", "addr", l.Addr().String())

	if err := s.httpServer.Serve(l); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("api: server error: %w", err)
	}

	return nil
}

// Shutdown gracefully stops the server with the provided context.
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down API server...")

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("api: shutdown error: %w", err)
	}

	s.logger.Info("API server stopped")

	return nil
}
