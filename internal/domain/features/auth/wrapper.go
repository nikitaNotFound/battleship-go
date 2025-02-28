package auth

import "github.com/nikitaNotFound/battleship/internal/domain/features"

type AuthFeatures struct {
	deps *features.FeaturesDeps
}

func NewAuthFeatures(deps *features.FeaturesDeps) *AuthFeatures {
	return &AuthFeatures{
		deps: deps,
	}
}
