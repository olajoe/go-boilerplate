package main

import (
	"fmt"
	"go-boilerplate/internal/config"
	"go-boilerplate/internal/repository/postgresql"
	"go-boilerplate/internal/rest"
	"go-boilerplate/internal/rest/middleware"
	"go-boilerplate/player"
	"log"
	"os"

	"github.com/labstack/echo/v4"
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
	log.Fatal(e.Start(address))
}
