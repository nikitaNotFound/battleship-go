package domain

import "github.com/nikitaNotFound/battleship/internal/domain/models"

type GameHistoryStorage interface {
	CreateGame(gameInfo models.GameInfo) error
	GetUserGameHistoryPaged(userID int64, page int, pageSize int) ([]models.GameInfo, error)
	GetGame(gameID int64) (models.GameInfo, error)
}

type UserStorage interface {
	CreateUser(user models.User) error
	GetUserByID(userID int64) (models.User, error)
	GetUserFriends(userID int64) ([]models.User, error)
	UpdateUserRank(userID int64, rankDelta int64) error
}
