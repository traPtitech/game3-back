-- +goose Up
ALTER TABLE game
    CHANGE COLUMN image image MEDIUMBLOB;

ALTER TABLE event
    CHANGE COLUMN image image MEDIUMBLOB;

-- +goose Down
ALTER TABLE game
    CHANGE COLUMN image image LONGBLOB;

ALTER TABLE event
    CHANGE COLUMN image image LONGBLOB;
