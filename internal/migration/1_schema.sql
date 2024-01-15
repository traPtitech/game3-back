-- +goose Up
-- +goose StatementBegin
CREATE TABLE sessions
(
    id            CHAR(36) PRIMARY KEY,
    redirect      VARCHAR(255) NOT NULL,
    access_token  VARCHAR(255),
    refresh_token VARCHAR(255),
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_in    INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sessions;
-- +goose StatementEnd
