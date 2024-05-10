package postgresql

import (
	"context"
	"go-boilerplate/domain"
)

type PlayerRepository struct {
}

func NewPlayerRepository() *PlayerRepository {
	return &PlayerRepository{}
}

func (m *PlayerRepository) GetByID(ctx context.Context, id string) (*domain.Player, error) {
	return &domain.Player{}, nil
}
