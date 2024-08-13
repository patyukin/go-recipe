package main

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-recipe/internal/client"
	"go-recipe/internal/config"
	"go-recipe/internal/db"
	"go-recipe/internal/dbconn"
	"go-recipe/internal/handler"
	"go-recipe/internal/server"
	"go-recipe/internal/server/router"
	"go-recipe/internal/usecase"
	"go-recipe/migrator"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("unable to load config: %v", err)
	}

	dbConn, err := dbconn.New(ctx, cfg)
	if err != nil {
		log.Fatal().Msgf("failed connecting to db: %v", err)
	}

	err = migrator.UpMigrations(ctx, dbConn)
	if err != nil {
		log.Fatal().Msgf("failed migrating db: %v", err)
	}

	dbClient := db.New(dbConn)
	clnt := client.NewClient(cfg)
	uc := usecase.New(dbClient, clnt)
	h := handler.New(uc)
	r := router.Init(ctx, h)

	srv := server.New(r)

	errCh := make(chan error)

	go func() {
		log.Info().Msgf("starting server on %d", cfg.HttpPort)
		if err = srv.Run(cfg); err != nil {
			log.Error().Msgf("failed starting server: %v", err)
			errCh <- err
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err = <-errCh:
		log.Error().Msgf("Failed to run, err: %v", err)
	case res := <-sigChan:
		if res == syscall.SIGINT || res == syscall.SIGTERM {
			log.Info().Msgf("Signal received")
		} else if res == syscall.SIGHUP {
			log.Info().Msgf("Signal received")
		}
	}

	log.Info().Msgf("Shutting Down")

	if err = srv.Shutdown(ctx); err != nil {
		log.Error().Msgf("failed server shutting down: %s", err.Error())
	}

	if err = dbClient.Close(); err != nil {
		log.Error().Msgf("failed db connection close: %s", err.Error())
	}
}
