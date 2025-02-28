package game

import "github.com/nikitaNotFound/battleship/internal/domain/features"

type GameFeatures struct {
	deps *features.FeaturesDeps
}

func NewGameFeatures(deps *features.FeaturesDeps) *GameFeatures {
	return &GameFeatures{
		deps: deps,
	}
}
