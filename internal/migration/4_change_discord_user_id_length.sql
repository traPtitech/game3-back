-- +goose Up
ALTER TABLE game MODIFY COLUMN discord_user_id VARCHAR(255) NOT NULL;

-- +goose Down
ALTER TABLE game MODIFY COLUMN discord_user_id CHAR(18) NOT NULL;
