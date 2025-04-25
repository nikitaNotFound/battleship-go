package core

import (
	"github.com/nikitaNotFound/battleship/internal/core/contracts"
	"github.com/nikitaNotFound/battleship/internal/core/features/auth"
	"github.com/nikitaNotFound/battleship/internal/core/features/game"
)

type Core struct {
	Auth    *auth.AuthFeatures
	Game    *game.GameFeatures
	Storage *contracts.Storage
}

func NewCore(storage *contracts.Storage) *Core {
	return &Core{
		Auth:    auth.NewAuthFeatures(storage),
		Game:    game.NewGameFeatures(storage),
		Storage: storage,
	}
}
