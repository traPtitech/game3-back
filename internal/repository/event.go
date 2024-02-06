package repository

import (
	"database/sql"
	"errors"
	"github.com/traPtitech/game3-back/openapi/models"
)

func selectEventWithoutImageQuery() string {
	return "SELECT event.slug, event.title, event.date , event.game_submission_period_start, event.game_submission_period_end FROM event "
}

func (r *Repository) GetEvents() ([]*models.Event, error) {
	events := []*models.Event{}
	query := selectEventWithoutImageQuery()
	if err := r.db.Select(&events, query); err != nil {
		return nil, err
	}

	return events, nil
}

func (r *Repository) PostEvent(event *models.PostEventRequest) (err error) {
	var imageData []byte
	if event.Image != nil {
		imageData, err = event.Image.Bytes()
		if err != nil {
			return err
		}
	}

	if _, err = r.db.Exec("INSERT INTO event (slug, title, date, game_submission_period_start, game_submission_period_end, image) VALUES (?, ?, ?, ?, ?, ?)", event.Slug, event.Title, event.Date, event.GameSubmissionPeriodStart, event.GameSubmissionPeriodEnd, imageData); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetCurrentEvent() (*models.Event, error) {
	var event models.Event

	futureQuery := selectEventWithoutImageQuery() + "WHERE date > NOW() ORDER BY date LIMIT 1"
	err := r.db.Get(&event, futureQuery)
	if err == nil {
		return &event, nil // 未来のイベントが見つかった
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	pastQuery := selectEventWithoutImageQuery() + "WHERE date <= NOW() ORDER BY date DESC LIMIT 1"
	err = r.db.Get(&event, pastQuery)
	if err == nil {
		return &event, nil // 過去のイベントが見つかった
	}

	return nil, err
}

func (r *Repository) GetEvent(eventSlug models.EventSlugInPath) (*models.Event, error) {
	event := &models.Event{}
	query := selectEventWithoutImageQuery() + "WHERE slug = ?"
	if err := r.db.Get(event, query, eventSlug); err != nil {
		return nil, err
	}

	return event, nil
}

func (r *Repository) PatchEvent(eventSlug models.EventSlugInPath, event *models.PatchEventRequest) error {
	return r.Patch("event", "slug", eventSlug, event)
}

func (r *Repository) DeleteEvent(eventSlug models.EventSlugInPath) error {
	_, err := r.db.Exec("DELETE FROM event WHERE slug = ?", eventSlug)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetEventImage(eventSlug models.EventSlugInPath) ([]byte, error) {
	file := []byte{}
	query := "SELECT image FROM event WHERE slug = ?"
	if err := r.db.Get(&file, query, eventSlug); err != nil {
		return nil, err
	}

	return file, nil
}
