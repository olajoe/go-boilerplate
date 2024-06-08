package ports

import "go-boilerplate/internal/domain"

type PlayerRepository interface {
	GetPlayerByID(id int) (*domain.Player, error)
}
