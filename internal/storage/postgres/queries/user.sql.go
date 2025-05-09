// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user.sql

package queries

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (username, password_hash, rank)
VALUES ($1, $2, $3)
`

type CreateUserParams struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
	Rank         int64  `json:"rank"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser, arg.Username, arg.PasswordHash, arg.Rank)
	return err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, password_hash, rank
FROM users
WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PasswordHash,
		&i.Rank,
	)
	return i, err
}

const getUserFriends = `-- name: GetUserFriends :many
SELECT u.id, u.username, u.password_hash, u.rank
FROM users u
WHERE u.id IN (
    SELECT uf.user2_id 
    FROM user_friends uf
    WHERE uf.user1_id = $1
) OR u.id IN (
    SELECT uf.user1_id 
    FROM user_friends uf
    WHERE uf.user2_id = $1
)
`

func (q *Queries) GetUserFriends(ctx context.Context, user1ID int64) ([]User, error) {
	rows, err := q.db.Query(ctx, getUserFriends, user1ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.PasswordHash,
			&i.Rank,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserRank = `-- name: UpdateUserRank :exec
UPDATE users
SET rank = rank + $2
WHERE id = $1
`

type UpdateUserRankParams struct {
	ID   int64 `json:"id"`
	Rank int64 `json:"rank"`
}

func (q *Queries) UpdateUserRank(ctx context.Context, arg UpdateUserRankParams) error {
	_, err := q.db.Exec(ctx, updateUserRank, arg.ID, arg.Rank)
	return err
}
