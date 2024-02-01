package domains

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID           *uuid.UUID `db:"id"`
	Redirect     *string    `db:"redirect"`
	AccessToken  *string    `db:"access_token"`
	RefreshToken *string    `db:"refresh_token"`
	ExpiresIn    *int       `db:"expires_in"`
	CreatedAt    *time.Time `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
}
