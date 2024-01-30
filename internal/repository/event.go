package repository

import "github.com/traPtitech/game3-back/internal/api/models"

func (r *Repository) GetEvents() ([]*models.Event, error) {
	events := []*models.Event{}
	if err := r.db.Select(&events, "SELECT * FROM event"); err != nil {
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

	if _, err = r.db.Exec("INSERT INTO event (slug, title, startAt, endAt, image) VALUES (?, ?, ?, ?, ?)", event.Slug, event.Title, event.GameSubmissionPeriodStart, event.GameSubmissionPeriodEnd, imageData); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetCurrentEvent() (*models.Event, error) {
	event := &models.Event{}
	if err := r.db.Get(event, "SELECT * FROM event WHERE startAt <= NOW() AND endAt >= NOW()"); err != nil {
		return nil, err
	}

	return event, nil
}

func (r *Repository) GetEvent(eventSlug models.EventSlugInPath) (*models.Event, error) {
	event := &models.Event{}
	if err := r.db.Get(event, "SELECT * FROM event WHERE slug = ?", eventSlug); err != nil {
		return nil, err
	}

	return event, nil
}
