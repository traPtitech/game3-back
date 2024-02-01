package domains

import (
	"github.com/google/uuid"
	"time"
)

/*
CREATE TABLE game
(
    id              CHAR(36) PRIMARY KEY,
    termId          CHAR(36) NOT NULL,
    discordUserId   CHAR(36) NOT NULL,
    creatorName     VARCHAR(255) NOT NULL,
    creatorPageUrl  VARCHAR(255),
    gamePageUrl     VARCHAR(255),
    title           VARCHAR(255) NOT NULL,
    description     TEXT,
    place           VARCHAR(255),
    icon            MEDIUMBLOB NOT NULL,
    image           MEDIUMBLOB,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (termId) REFERENCES term(id),
    FOREIGN KEY (discordUserId) REFERENCES session(id)
);
*/

type Game struct {
	ID             *uuid.UUID `db:"id" json:"id,omitempty"`
	TermID         *uuid.UUID `db:"termId" json:"term_id,omitempty"`
	DiscordUserID  *string    `db:"discordUserId" json:"discord_user_id,omitempty"`
	CreatorName    *string    `db:"creatorName" json:"creator_name,omitempty"`
	CreatorPageURL *string    `db:"creatorPageUrl" json:"creator_page_url,omitempty"`
	GamePageURL    *string    `db:"gamePageUrl" json:"game_page_url,omitempty"`
	Title          *string    `db:"title" json:"title,omitempty"`
	Description    *string    `db:"description" json:"description,omitempty"`
	Place          *string    `db:"place" json:"place,omitempty"`
	Icon           *[]byte    `db:"icon" json:"icon,omitempty"`
	Image          *[]byte    `db:"image" json:"image,omitempty"`
	CreatedAt      *time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
