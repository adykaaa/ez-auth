package http

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
)

type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Server struct {
	l               *zerolog.Logger
	s               *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func NewServer(r Router, addr string, l *zerolog.Logger) (*Server, error) {
	s := &Server{
		s: &http.Server{
			Handler: r,
			Addr:    addr,
		},
		notify:          make(chan error, 1),
		shutdownTimeout: 5 * time.Second,
		l:               l,
	}

	if addr == "" {
		s.l.Error().Msg("server address is empty")
		return nil, errors.New("server address cannot be empty")
	}

	s.Start()
	return s, nil
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.s.ListenAndServe()
		close(s.notify)
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-interrupt:
		s.l.Info().Msgf("Server run interrupted by OS signal %s", sig.String())
	case err := <-s.notify:
		s.l.Error().Msgf("error during server connection %v", err)
	}
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.s.Shutdown(ctx)
}
