package auth

import "github.com/nikitaNotFound/battleship/internal/core/contracts"

type AuthFeatures struct {
	storage *contracts.Storage
}

func NewAuthFeatures(storage *contracts.Storage) *AuthFeatures {
	return &AuthFeatures{
		storage: storage,
	}
}
