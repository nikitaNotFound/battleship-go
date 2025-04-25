// users.go
package server

import (
	"github.com/labstack/echo/v4"
	"github.com/nikitaNotFound/battleship/internal/webserver/apigen"
)

func (s *BattleshipServer) CreateUser(ctx echo.Context) error {
	return nil
}

func (s *BattleshipServer) GetUserById(ctx echo.Context, userId int64) error {
	return nil
}

func (s *BattleshipServer) GetUserFriends(ctx echo.Context, userId int64) error {
	return nil
}

func (s *BattleshipServer) GetUserGameHistory(ctx echo.Context, userId int64, params apigen.GetUserGameHistoryParams) error {
	return nil
}

func (s *BattleshipServer) UpdateUserRank(ctx echo.Context, userId int64) error {
	return nil
}
