package domains

import (
	"github.com/google/uuid"
	"time"
)

/*
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
*/

type Session struct {
	ID           *uuid.UUID `db:"id"`
	Redirect     *string    `db:"redirect"`
	AccessToken  *string    `db:"access_token"`
	RefreshToken *string    `db:"refresh_token"`
	ExpiresIn    *int       `db:"expires_in"`
	CreatedAt    *time.Time `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
}
