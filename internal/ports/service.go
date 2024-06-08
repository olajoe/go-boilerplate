package ports

import "go-boilerplate/internal/domain"

type PlayerService interface {
	GetPlayerByID(id int) (*domain.Player, error)
}
