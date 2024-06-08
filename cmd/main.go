package main

import (
	"context"
	"fmt"
	"go-boilerplate/internal/adapter/db"
	_http "go-boilerplate/internal/adapter/http"
	"go-boilerplate/internal/config"
	"go-boilerplate/internal/usecase"
	"go-boilerplate/internal/utils"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.NewConfiguration()

	playeRepo := db.NewPlayerRepositoryDB()

	playerUseCase := usecase.NewPlayerUseCase(playeRepo)

	handler := _http.NewPlayerHandler(playerUseCase)

	router := mux.NewRouter()

	router.HandleFunc("/healthz", utils.HealthCheckHandler).Methods(http.MethodGet)

	router.HandleFunc("/players/:id", handler.GetPlayerByID).Methods(http.MethodGet)

	srv := http.Server{
		Addr:         fmt.Sprintf("%s:%s", "localhost", cfg.Port),
		Handler:      router,
		WriteTimeout: time.Second * time.Duration(cfg.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(cfg.ReadTimeout),
		IdleTimeout:  time.Second * time.Duration(cfg.IdleTimeout),
	}

	go func() {
		slog.Info(fmt.Sprintf("Server listening on %v", cfg.Port))
		if err := srv.ListenAndServe(); err != nil {
			slog.Error(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(cfg.Timeout))
	defer cancel()

	srv.Shutdown(ctx)

	slog.Info("shutting down")
	os.Exit(0)
}
