-- name: CreateGame :exec
INSERT INTO games (
    player1_id, player2_id, player1_rank_change, 
    player2_rank_change, created_at, finished_at
) VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetUserGameHistoryPaged :many
SELECT *
FROM games
WHERE player1_id = $1 OR player2_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetGameById :one
SELECT *
FROM games
WHERE id = $1; 