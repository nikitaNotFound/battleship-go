package models

import "time"

type Leaderboard struct {
	CreatedAt time.Time
	Users     []User
}
