package main

import (
	"database/sql"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
	tadokumiddleware "github.com/tadoku/tadoku/services/common/middleware"
	"github.com/tadoku/tadoku/services/common/storage/memory"
	"github.com/tadoku/tadoku/services/content-api/domain/pagecommand"
	"github.com/tadoku/tadoku/services/content-api/domain/pagequery"
	"github.com/tadoku/tadoku/services/content-api/domain/postcommand"
	"github.com/tadoku/tadoku/services/content-api/http/rest"
	"github.com/tadoku/tadoku/services/content-api/http/rest/openapi"
	"github.com/tadoku/tadoku/services/content-api/storage/postgres"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
)

type Config struct {
	PostgresURL string `validate:"required" envconfig:"postgres_url"`
	Port        int64  `validate:"required"`
	JWKS        string `validate:"required"`
}

func main() {
	cfg := Config{}
	envconfig.Process("API", &cfg)

	validate := validator.New()
	err := validate.Struct(cfg)
	if err != nil {
		panic(fmt.Errorf("could not configure server: %w", err))
	}

	psql, err := sql.Open("pgx", cfg.PostgresURL)
	if err != nil {
		panic(err)
	}

	pageRepository := postgres.NewPageRepository(psql)
	postRepository := postgres.NewPostRepository(psql)
	roleRepository := memory.NewRoleRepository("/etc/tadoku/permissions/roles.yaml")

	e := echo.New()
	e.Use(tadokumiddleware.Logger([]string{"/ping"}))
	e.Use(tadokumiddleware.SessionJWT(cfg.JWKS))
	e.Use(tadokumiddleware.Session(roleRepository))

	pageCommandService := pagecommand.NewService(pageRepository)
	postCommandService := postcommand.NewService(postRepository)

	pageQueryService := pagequery.NewService(pageRepository)

	server := rest.NewServer(
		pageCommandService,
		postCommandService,

		pageQueryService,
	)

	openapi.RegisterHandlersWithBaseURL(e, server, "")

	fmt.Printf("content-api is now available at: http://localhost:%d/v2\n", cfg.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", cfg.Port)))
}
