package app

import (
	"context"
	"fmt"
	"github.com/artarts36/lowboard/registry/internal/port/handlers"
	"github.com/artarts36/lowboard/registry/internal/repository"
	"log/slog"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // sql driver

	"github.com/artarts36/lowboard/registry/internal/port/generated/api"
	"github.com/artarts36/lowboard/registry/internal/port/middlewares"
)

type Application struct {
	config *Config

	server *http.Server

	repositories *repository.Group

	db *sqlx.DB
}

func NewApplication(cfg *Config, db *sqlx.DB) (*Application, error) {
	s := &Application{
		config: cfg,
	}

	s.db = db

	if err := s.initInfrastructure(); err != nil {
		return nil, fmt.Errorf("init infrastructure: %w", err)
	}

	srv, err := api.NewServer(handlers.NewService(s.repositories))
	if err != nil {
		return nil, fmt.Errorf("init http handlers: %w", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", s.buildHandler(srv))

	s.server = &http.Server{
		Addr:        s.config.HTTP.Addr,
		Handler:     mux,
		ReadTimeout: 30 * time.Second,
	}

	return s, nil
}

func (s *Application) Run() error {
	slog.Info("[app] listen http", slog.String("addr", s.config.HTTP.Addr))

	return s.server.ListenAndServe()
}

func (s *Application) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Application) buildHandler(next *api.Server) http.Handler {
	return middlewares.NewRecovery(
		middlewares.CORS(next, s.config.HTTP.Clients),
	)
}

func (s *Application) initInfrastructure() error {
	s.repositories = repository.NewGroup(s.db)

	return nil
}
