package services

import "github.com/nikitaNotFound/battleship/internal/storage"

type GameService struct {
	storage *storage.BattleshipStorage
}

func NewGameService(storage *storage.BattleshipStorage) *GameService {
	return &GameService{
		storage: storage,
	}
}
