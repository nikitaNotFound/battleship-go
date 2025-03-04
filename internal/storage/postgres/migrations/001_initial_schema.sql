-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    rank BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE games (
    id BIGSERIAL PRIMARY KEY,
    player1_id BIGINT NOT NULL REFERENCES users(id),
    player2_id BIGINT NOT NULL REFERENCES users(id),
    player1_rank_change BIGINT NOT NULL,
    player2_rank_change BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    finished_at TIMESTAMP NOT NULL,
    CONSTRAINT different_players CHECK (player1_id != player2_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS games;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd 