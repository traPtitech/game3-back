-- +goose Up
-- +goose StatementBegin
CREATE TABLE session
(
    id            CHAR(36) PRIMARY KEY,
    redirect      VARCHAR(255) NOT NULL,
    access_token  VARCHAR(255),
    refresh_token VARCHAR(255),
    expires_in    INT,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE event
(
    slug                         VARCHAR(255) PRIMARY KEY,
    title                        VARCHAR(255) NOT NULL,
    date                         TIMESTAMP    NOT NULL,
    game_submission_period_start TIMESTAMP    NOT NULL,
    game_submission_period_end   TIMESTAMP    NOT NULL,
    image                        LONGBLOB,
    created_at                   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at                   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE term
(
    id         CHAR(36) PRIMARY KEY,
    event_slug VARCHAR(255) NOT NULL,
    is_default BOOLEAN      NOT NULL,
    start_at   TIMESTAMP    NOT NULL,
    end_at     TIMESTAMP    NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (event_slug) REFERENCES event (slug)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE game
(
    id               CHAR(36) PRIMARY KEY,
    term_id          CHAR(36)     NOT NULL,
    discord_user_id  CHAR(18)     NOT NULL,
    creator_name     VARCHAR(255) NOT NULL,
    creator_page_url VARCHAR(255),
    game_page_url    VARCHAR(255),
    title            VARCHAR(255) NOT NULL,
    description      TEXT,
    place            VARCHAR(255),
    icon             MEDIUMBLOB   NOT NULL,
    image            LONGBLOB,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (term_id) REFERENCES term (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS session;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS game;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS term;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS event;
-- +goose StatementEnd
