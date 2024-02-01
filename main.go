package main

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/game3-back/internal/handler"
	"github.com/traPtitech/game3-back/internal/migration"
	"github.com/traPtitech/game3-back/internal/pkg/config"
	"github.com/traPtitech/game3-back/internal/pkg/util"
	"github.com/traPtitech/game3-back/internal/repository"
	"github.com/traPtitech/game3-back/openapi"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	swagger, err := openapi.GetSwagger()
	if err != nil {
		fmt.Printf("Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	baseUrl := "/api"
	swagger.Servers = openapi3.Servers{&openapi3.Server{URL: baseUrl}}
	// middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	//e.Use(oapimiddleware.OapiRequestValidator(swagger))

	// connect to database
	db, err := sqlx.Connect("mysql", config.MySQL().FormatDSN())
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()
	sqlx.NameMapper = util.ToSnakeCase

	// migrate tables
	if err := migration.MigrateTables(db.DB); err != nil {
		e.Logger.Fatal(err)
	}

	// setup repository
	repo := repository.New(db)

	// setup routes
	h := handler.New(repo)
	openapi.RegisterHandlersWithBaseURL(e, h, baseUrl)

	e.Logger.Fatal(e.Start(config.AppAddr()))
}
