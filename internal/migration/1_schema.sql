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
    slug       VARCHAR(255) PRIMARY KEY,
    title      VARCHAR(255) NOT NULL,
    startAt    TIMESTAMP    NOT NULL,
    endAt      TIMESTAMP    NOT NULL,
    image      LONGBLOB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE term
(
    id         CHAR(36) PRIMARY KEY,
    eventSlug  VARCHAR(255) NOT NULL,
    isDefault  BOOLEAN      NOT NULL,
    startAt    TIMESTAMP    NOT NULL,
    endAt      TIMESTAMP    NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (eventSlug) REFERENCES event (slug)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE game
(
    id             CHAR(36) PRIMARY KEY,
    termId         CHAR(36)     NOT NULL,
    discordUserId  CHAR(36)     NOT NULL,
    creatorName    VARCHAR(255) NOT NULL,
    creatorPageUrl VARCHAR(255),
    gamePageUrl    VARCHAR(255),
    title          VARCHAR(255) NOT NULL,
    description    TEXT,
    place          VARCHAR(255),
    icon           MEDIUMBLOB   NOT NULL,
    image          LONGBLOB,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (termId) REFERENCES term (id),
    FOREIGN KEY (discordUserId) REFERENCES session (id)
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
