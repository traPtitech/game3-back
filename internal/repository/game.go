package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/traPtitech/game3-back/internal/api/models"
	"github.com/traPtitech/game3-back/internal/domains"
	"strings"
)

func selectGameWithoutImagesQuery() string {
	return "SELECT game.id, game.termId, game.discordUserId, game.creatorName, game.creatorPageUrl, game.gamePageUrl, game.title, game.description, game.place, game.created_at, game.updated_at FROM game "
}

func (r *Repository) GetGames(params models.GetGamesParams) ([]*domains.Game, error) {
	games := []*domains.Game{}
	query := selectGameWithoutImagesQuery()
	whereClauses := []string{}
	args := []interface{}{}

	if params.TermId != nil {
		whereClauses = append(whereClauses, "game.termId = ?")
		args = append(args, params.TermId)
	}
	if params.EventId != nil {
		whereClauses = append(whereClauses, "game.eventId = ?")
		args = append(args, params.EventId)
	}
	if params.UserId != nil {
		whereClauses = append(whereClauses, "game.discordUserId = ?")
		args = append(args, params.UserId)
	}
	if params.Include != nil && *params.Include == "unpublished" {
		query += "JOIN term ON game.termId = term.id "
		whereClauses = append(whereClauses, "term.isDefault = TRUE")
	}

	if len(whereClauses) > 0 {
		query += "WHERE " + strings.Join(whereClauses, " AND ")
	}

	if err := r.db.Select(&games, query, args...); err != nil {
		return nil, err
	}

	return games, nil
}

func (r *Repository) PostGame(game *models.PostGameRequest) (*uuid.UUID, error) {
	newGameID := uuid.New()
	// TODO default termId and Session

	iconData, err := game.Icon.Bytes()
	if err != nil {
		return nil, err
	}
	fmt.Println(game.Icon.Filename())
	var imageData []byte
	if game.Image != nil {
		imageData, err = game.Image.Bytes()
		if err != nil {
			return nil, err
		}
	}
	_, err = r.db.Exec("INSERT INTO game (id, termId, discordUserId, creatorName, creatorPageUrl, gamePageUrl, title, description, icon, image) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", newGameID, uuid.UUID{}, uuid.UUID{}, game.CreatorName, game.CreatorPageUrl, game.GamePageUrl, game.Title, game.Description, iconData, imageData)
	if err != nil {
		return nil, err
	}
	return &newGameID, nil
}

func (r *Repository) GetGame(gameID uuid.UUID) (*domains.Game, error) {
	game := &domains.Game{}
	query := selectGameWithoutImagesQuery() + "WHERE id = ?"
	if err := r.db.Get(game, query, gameID); err != nil {
		return nil, err
	}

	return game, nil
}
