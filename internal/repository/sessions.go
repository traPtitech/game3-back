package repository

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID           *uuid.UUID `db:"id"`
	Redirect     *string    `db:"redirect"`
	AccessToken  *string    `db:"access_token"`
	RefreshToken *string    `db:"refresh_token"`
	CreatedAt    *time.Time `db:"created_at"`
	ExpiresIn    *int       `db:"expires_in"`
}

func (r *Repository) CreateSession(session *Session) error {
	_, err := r.db.Exec("INSERT INTO sessions (id, redirect) VALUES (?, ?)", *session.ID, *session.Redirect)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateSession(session *Session) error {
	_, err := r.db.Exec("UPDATE sessions SET access_token = ?, refresh_token = ?, expires_in = ? WHERE id = ?", *session.AccessToken, *session.RefreshToken, *session.ExpiresIn, *session.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetSession(sessionID string) (*Session, error) {
	session := &Session{}
	if err := r.db.Get(session, "SELECT * FROM sessions WHERE id = ?", sessionID); err != nil {
		return nil, err
	}

	return session, nil
}
