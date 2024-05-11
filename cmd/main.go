package main

import (
	"context"
	"fmt"
	"go-boilerplate/internal/config"
	"go-boilerplate/internal/repository/postgresql"
	"go-boilerplate/internal/rest"
	"go-boilerplate/internal/rest/middleware"
	"go-boilerplate/player"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const (
	defaultTimeout = 30
	defaultAddress = ":9090"
)

func main() {
	fmt.Println("Hello world")

	cfg := config.NewConfiguration()
	_ = cfg

	e := echo.New()
	e.Use(middleware.CORS)

	// Prepare Repository
	playerRepo := postgresql.NewPlayerRepository()

	// Build service Layer
	playerServ := player.NewPlayerService(playerRepo)
	rest.NewPlayerHandler(e, playerServ)

	// Start Server
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(address); err != nil && err != http.ErrServerClosed {
			logrus.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logrus.Fatal(err)
	}

}
