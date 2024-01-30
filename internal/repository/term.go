package repository

import (
	"github.com/google/uuid"
	"github.com/traPtitech/game3-back/internal/api/models"
)

func (r *Repository) GetTerms() ([]*models.Term, error) {
	terms := []*models.Term{}
	if err := r.db.Select(&terms, "SELECT * FROM term"); err != nil {
		return nil, err
	}

	return terms, nil
}

func (r *Repository) PostTerm(newTermID uuid.UUID, req *models.PostTermRequest) error {
	if _, err := r.db.Exec("INSERT INTO term (id, eventSlug, isDefault, startAt, endAt) VALUES (?, ?, ?, ?, ?)", newTermID, req.EventSlug, false, req.StartAt, req.EndAt); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetTerm(termID uuid.UUID) (*models.Term, error) {
	term := &models.Term{}
	if err := r.db.Get(term, "SELECT * FROM term WHERE id = ?", termID); err != nil {
		return nil, err
	}

	return term, nil
}

func (r *Repository) PatchTerm(termID uuid.UUID, req *models.PatchTermRequest) error {
	if _, err := r.db.Exec("UPDATE term SET eventSlug = ?, isDefault = ?, startAt = ?, endAt = ? WHERE id = ?", req.EventSlug, req.IsDefault, req.StartAt, req.EndAt, termID); err != nil {
		return err
	}

	return nil
}
