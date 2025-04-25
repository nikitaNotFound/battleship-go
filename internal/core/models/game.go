package models

import "time"

type GameInfo struct {
	ID                int64
	Player1ID         int64
	Player2ID         int64
	Player1RankChange int64
	Player2RankChange int64
	CreatedAt         time.Time
	FinishedAt        time.Time
}
