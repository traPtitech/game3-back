package repository

import (
	"github.com/traPtitech/game3-back/internal/api/models"
)

func selectEventWithoutImageQuery() string {
	return "SELECT event.slug, event.title, event.game_submission_period_start, event.game_submission_period_end FROM event "
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

	if _, err = r.db.Exec("INSERT INTO event (slug, title, game_submission_period_start, game_submission_period_end, image) VALUES (?, ?, ?, ?, ?)", event.Slug, event.Title, event.GameSubmissionPeriodStart, event.GameSubmissionPeriodEnd, imageData); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetCurrentEvent() (*models.Event, error) {
	event := &models.Event{}
	query := selectEventWithoutImageQuery() + "WHERE game_submission_period_start <= NOW() AND game_submission_period_end >= NOW()"
	if err := r.db.Get(event, query); err != nil {
		return nil, err
	}

	return event, nil
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

func (r *Repository) GetEventTerms(eventSlug models.EventSlugInPath) ([]*models.Term, error) {
	terms := []*models.Term{}
	query := "SELECT term.* FROM term JOIN event ON term.event_slug = event.slug WHERE event.slug = ?"
	if err := r.db.Select(&terms, query, eventSlug); err != nil {
		return nil, err
	}

	return terms, nil
}

func (r *Repository) GetEventGames(eventSlug models.EventSlugInPath) ([]*models.Game, error) {
	games := []*models.Game{}
	query := selectGameWithoutImagesQuery() + `
FROM game
JOIN term ON game.term_id = term.id
JOIN event ON term.event_slug = event.slug
WHERE event.slug = ?`
	if err := r.db.Select(&games, query, eventSlug); err != nil {
		return nil, err
	}

	return games, nil
}

func (r *Repository) GetEventImage(eventSlug models.EventSlugInPath) ([]byte, error) {
	file := []byte{}
	query := "SELECT image FROM event WHERE slug = ?"
	if err := r.db.Get(&file, query, eventSlug); err != nil {
		return nil, err
	}

	return file, nil
}
