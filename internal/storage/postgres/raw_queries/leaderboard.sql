-- name: GetLeaderboard :many
SELECT id, username, rank
FROM users
ORDER BY rank DESC
LIMIT $1 OFFSET $2;