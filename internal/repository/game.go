package repository

import (
	"github.com/google/uuid"
	"github.com/traPtitech/game3-back/internal/domain"
	"github.com/traPtitech/game3-back/openapi/models"
	"strings"
)

func selectGameWithoutImagesQuery() string {
	return "SELECT game.id, game.term_id, game.discord_user_id, game.is_published, game.creator_name, game.creator_page_url, game.game_page_url, game.title, game.description, game.place FROM game "
}

func (r *Repository) GetGames(params models.GetGamesParams) ([]*models.Game, error) {
	games := []*models.Game{}
	query := selectGameWithoutImagesQuery()
	whereClauses := []string{}
	args := []interface{}{}

	// term_idでフィルタ
	if params.TermId != nil {
		whereClauses = append(whereClauses, "game.term_id = ?")
		args = append(args, params.TermId)
	}

	// event_slugでフィルタ、eventとtermテーブルを結合
	if params.EventSlug != nil {
		query += `JOIN term ON game.term_id = term.id 
                  JOIN event ON term.event_slug = event.slug `
		whereClauses = append(whereClauses, "event.slug = ?")
		args = append(args, params.EventSlug)
	}

	// discordUserIdでフィルタ
	if params.UserId != nil {
		whereClauses = append(whereClauses, "game.discord_user_id = ?")
		args = append(args, params.UserId)
	}

	if params.IncludeUnpublished == nil || !*params.IncludeUnpublished {
		whereClauses = append(whereClauses, "game.is_published = TRUE")
	}

	// WHERE句の組み立て
	if len(whereClauses) > 0 {
		query += "WHERE " + strings.Join(whereClauses, " AND ")
	}

	if err := r.db.Select(&games, query, args...); err != nil {
		return nil, err
	}

	return games, nil
}

func (r *Repository) PostGame(newGameID uuid.UUID, termID uuid.UUID, userID string, game *models.PostGameRequest) error {
	iconData, err := game.Icon.Bytes()
	if err != nil {
		return err
	}

	var imageData []byte
	if game.Image != nil {
		imageData, err = game.Image.Bytes()
		if err != nil {
			return err
		}
	}
	_, err = r.db.Exec("INSERT INTO game (id, term_id, discord_user_id, creator_name, creator_page_url, game_page_url, title, description, icon, image) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", newGameID, termID, userID, game.CreatorName, game.CreatorPageUrl, game.GamePageUrl, game.Title, game.Description, iconData, imageData)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetGame(gameID uuid.UUID) (*models.Game, error) {
	game := &models.Game{}
	query := selectGameWithoutImagesQuery() + "WHERE id = ?"
	if err := r.db.Get(game, query, gameID); err != nil {
		return nil, err
	}

	return game, nil
}

func (r *Repository) PatchGame(gameID uuid.UUID, game *models.PatchGameRequest) error {
	return r.Patch("game", "id", gameID, game)
}

func (r *Repository) DeleteGame(gameID uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM game WHERE id = ?", gameID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetGameIcon(gameID uuid.UUID) (*domain.Icon, error) {
	icon := &domain.Icon{}
	if err := r.db.Get(icon, "SELECT icon, updated_at FROM game WHERE id = ?", gameID); err != nil {
		return nil, err
	}

	return icon, nil
}

func (r *Repository) GetGameImage(gameID uuid.UUID) (*domain.Image, error) {
	image := &domain.Image{}
	if err := r.db.Get(image, "SELECT image, updated_at FROM game WHERE id = ?", gameID); err != nil {
		return nil, err
	}

	return image, nil
}
