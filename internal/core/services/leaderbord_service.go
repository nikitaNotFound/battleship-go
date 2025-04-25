package services

import (
	"github.com/nikitaNotFound/battleship/internal/domain/models"
	"github.com/nikitaNotFound/battleship/internal/storage"
)

type LeaderboardService struct {
	storage *storage.BattleshipStorage
}

func NewLeaderboardService(storage *storage.BattleshipStorage) *LeaderboardService {
	return &LeaderboardService{
		storage: storage,
	}
}

func (s *LeaderboardService) GetLeaderboard(page int, pageSize int) (*models.Leaderboard, error) {
	return nil, nil
}
