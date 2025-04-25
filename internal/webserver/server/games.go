package server

import (
	"github.com/labstack/echo/v4"
	"github.com/nikitaNotFound/battleship/internal/webserver/apigen"
)

func (s *BattleshipServer) CreateGame(ctx echo.Context) error {
	return nil
}

func (s *BattleshipServer) GetGame(ctx echo.Context, gameId int64) error {
	return nil
}

func (s *BattleshipServer) GetLeaderboard(ctx echo.Context, params apigen.GetLeaderboardParams) error {
	return nil
}

func (s *BattleshipServer) SetLeaderboard(ctx echo.Context, params apigen.SetLeaderboardParams) error {
	return nil
}
