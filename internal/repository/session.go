package repository

import (
	"github.com/google/uuid"
	"github.com/traPtitech/game3-back/internal/domains"
)

func (r *Repository) CreateSession(session *domains.Session) error {
	_, err := r.db.Exec("INSERT INTO session (id, redirect) VALUES (?, ?)", *session.ID, *session.Redirect)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateSession(session *domains.Session) error {
	_, err := r.db.Exec("UPDATE session SET access_token = ?, refresh_token = ?, expires_in = ? WHERE id = ?", *session.AccessToken, *session.RefreshToken, *session.ExpiresIn, *session.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetSession(sessionID *uuid.UUID) (*domains.Session, error) {
	session := &domains.Session{}
	if err := r.db.Get(session, "SELECT * FROM session WHERE id = ?", sessionID); err != nil {
		return nil, err
	}

	return session, nil
}
