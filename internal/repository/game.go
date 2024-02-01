package repository

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/game3-back/openapi/models"
	"net/http"
	"strings"
)

func selectGameWithoutImagesQuery() string {
	return "SELECT game.id, game.term_id, game.discord_user_id, game.creator_name, game.creator_page_url, game.game_page_url, game.title, game.description, game.place FROM game "
}

func (r *Repository) GetGames(params models.GetGamesParams) ([]*models.Game, error) {
	games := []*models.Game{}
	query := selectGameWithoutImagesQuery()
	whereClauses := []string{}
	args := []interface{}{}
	joinedTerm := false

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
		joinedTerm = true
	}

	// discordUserIdでフィルタ
	if params.UserId != nil {
		whereClauses = append(whereClauses, "game.discord_user_id = ?")
		args = append(args, params.UserId)
	}

	if !joinedTerm {
		// termテーブルをJOINしていない場合、ここでJOIN
		query += `JOIN term ON game.term_id = term.id `
	}
	if params.Include != nil {
		if *params.Include == "unpublished" {
			whereClauses = append(whereClauses, "term.is_default = TRUE")
		} else {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "includeに指定できる値は'unpublished'のみ")
		}
	} else {
		whereClauses = append(whereClauses, "term.is_default = FALSE")
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

func (r *Repository) GetGameIcon(gameID uuid.UUID) ([]byte, error) {
	icon := []byte{}
	if err := r.db.Get(&icon, "SELECT icon FROM game WHERE id = ?", gameID); err != nil {
		return nil, err
	}

	return icon, nil
}

func (r *Repository) GetGameImage(gameID uuid.UUID) ([]byte, error) {
	image := []byte{}
	if err := r.db.Get(&image, "SELECT icon FROM game WHERE id = ?", gameID); err != nil {
		return nil, err
	}

	return image, nil
}
