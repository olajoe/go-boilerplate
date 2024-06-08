package usecase

import (
	"go-boilerplate/internal/domain"
	"go-boilerplate/internal/ports"
)

type PlayerUseCase struct {
	repo ports.PlayerRepository
}

func NewPlayerUseCase(repo ports.PlayerRepository) ports.PlayerService {
	return &PlayerUseCase{repo: repo}
}

func (uc *PlayerUseCase) GetPlayerByID(id int) (*domain.Player, error) {
	return uc.repo.GetPlayerByID(id)
}
