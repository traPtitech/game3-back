package main

import (
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/game3-back/internal/handler"
	"github.com/traPtitech/game3-back/internal/migration"
	"github.com/traPtitech/game3-back/internal/pkg/config"
	"github.com/traPtitech/game3-back/internal/repository"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// connect to database
	db, err := sqlx.Connect("mysql", config.MySQL().FormatDSN())
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	// migrate tables
	if err := migration.MigrateTables(db.DB); err != nil {
		e.Logger.Fatal(err)
	}

	// setup repository
	repo := repository.New(db)

	// setup routes
	h := handler.New(repo)
	v1API := e.Group("/api/v1")
	h.SetupRoutes(v1API)

	e.Logger.Fatal(e.Start(config.AppAddr()))
}
