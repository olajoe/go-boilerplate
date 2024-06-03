package rest

import (
	"context"
	"go-boilerplate/domain"
	"go-boilerplate/errs"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PlayerService interface {
	GetByID(ctx context.Context, id string) (*domain.Player, error)
}

type PlayerHandler struct {
	Service PlayerService
}

func NewPlayerHandler(e *echo.Echo, svc PlayerService) {
	handler := &PlayerHandler{
		Service: svc,
	}

	e.GET("/players/:id", handler.GetByID)
}

func (p *PlayerHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()

	player, err := p.Service.GetByID(ctx, id)
	if err != nil {
		return c.JSON(errs.GetStatusCode(err), errs.NewResponseError(err.Error()))
	}

	return c.JSON(http.StatusOK, player)
}
