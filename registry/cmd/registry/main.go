package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/artarts36/lowboard/registry/internal/port/app"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/pressly/goose"
)

func main() {
	slog.Info("read config")

	cfg, err := app.InitConfig("LOWBOARD_")
	if err != nil {
		slog.Info("failed to read config", slog.Any("err", err))
		os.Exit(1)
	}

	db, err := initDB(cfg)
	if err != nil {
		slog.Info("failed to init database", slog.Any("err", err))
		os.Exit(1)
	}

	if err = runMigrations(db); err != nil {
		slog.Info("failed to run migrations", slog.Any("err", err))
		os.Exit(1)
	}

	appServer, err := app.NewApplication(cfg, db)
	if err != nil {
		slog.Error("create application")
		os.Exit(1)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	wg := &sync.WaitGroup{}

	go func(s *app.Application) {
		slog.Info("run application")
		wg.Add(1)

		if serveErr := s.Run(); serveErr != nil {
			slog.
				With(slog.Any("err", serveErr)).
				Error("failed to serve http")
		}

		wg.Done()
	}(appServer)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)

	<-done
	if appServeErr := appServer.Shutdown(ctx); appServeErr != nil {
		slog.
			With(slog.String("err", appServeErr.Error())).
			Error("http server shutdown failed")
	}

	slog.Info("Cancelling root context")
	cancel()

	wg.Wait()
}

func initDB(config *app.Config) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	if config.DB.Driver == "sqlite" {
		db, err = sqlx.Open("sqlite3", config.DB.DSN)
		if err != nil {
			return nil, err
		}

		_, ferr := os.Open(config.DB.DSN)
		if ferr != nil && errors.Is(err, os.ErrNotExist) {
			_, cerr := os.Create(config.DB.DSN)
			if cerr != nil {
				return nil, fmt.Errorf("create db file: %w", err)
			}
		}

		_ = goose.SetDialect("sqlite3")
	}

	return db, nil
}

func runMigrations(db *sqlx.DB) error {
	return goose.Up(db.DB, "/app/migrations")
}
