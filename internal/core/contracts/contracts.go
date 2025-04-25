package contracts

import (
	"time"

	"github.com/nikitaNotFound/battleship/internal/core/models"
)

type Storage interface {
	// Game
	CreateGame(gameInfo models.GameInfo) error
	GetUserGameHistoryPaged(userID int64, page int, pageSize int) ([]models.GameInfo, error)
	GetGame(gameID int64) (models.GameInfo, error)

	// User
	CreateUser(user models.User) error
	GetUserByID(userID int64) (models.User, error)
	GetUserFriends(userID int64) ([]models.User, error)
	UpdateUserRank(userID int64, rankDelta int64) error

	// Leaderboard
	GetLeaderboard(page int, pageSize int) ([]models.User, error)
	SetLeaderboard(leaderboard models.Leaderboard, aliveFor time.Duration) error
}
