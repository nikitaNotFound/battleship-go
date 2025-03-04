-- name: CreateUser :exec
INSERT INTO users (username, password_hash, rank)
VALUES ($1, $2, $3);

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserFriends :many
SELECT u.*
FROM users u
WHERE u.id IN (
    SELECT uf.user2_id 
    FROM user_friends uf
    WHERE uf.user1_id = $1
) OR u.id IN (
    SELECT uf.user1_id 
    FROM user_friends uf
    WHERE uf.user2_id = $1
);

-- name: UpdateUserRank :exec
UPDATE users
SET rank = rank + $2
WHERE id = $1; 