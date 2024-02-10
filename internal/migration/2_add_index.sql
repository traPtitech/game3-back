-- +goose Up
ALTER TABLE event ADD INDEX idx_event_date (date);
ALTER TABLE game ADD INDEX idx_game_discord_user_id (discord_user_id);
ALTER TABLE game ADD INDEX idx_game_is_published (is_published);

-- +goose Down
ALTER TABLE event DROP INDEX idx_event_date;
ALTER TABLE game DROP INDEX idx_game_discord_user_id;
ALTER TABLE game DROP INDEX idx_game_is_published;
