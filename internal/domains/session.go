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
	ID           *uuid.UUID
	Redirect     *string
	AccessToken  *string
	RefreshToken *string
	ExpiresIn    *int
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
