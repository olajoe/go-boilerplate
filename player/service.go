package player

import (
	"context"
	"go-boilerplate/domain"
)

type PlayerRepository interface {
	GetByID(ctx context.Context, id string) (*domain.Player, error)
}

type PlayerService struct {
	playerRepo PlayerRepository
}

func NewPlayerService(p PlayerRepository) *PlayerService {
	return &PlayerService{
		playerRepo: p,
	}
}

func (p *PlayerService) GetByID(ctx context.Context, id string) (*domain.Player, error) {
	resPlayer, err := p.playerRepo.GetByID(ctx, id)
	if err != nil {
		return &domain.Player{}, err
	}
	return resPlayer, nil
}
