package webserver

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nikitaNotFound/battleship/internal/webserver/apigen"
	"github.com/nikitaNotFound/battleship/internal/webserver/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Run(ctx context.Context) error {
	cfg, err := GetWebServerConfig()
	if err != nil {
		return err
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	server := server.BattleshipServer{}

	e := echo.New()

	// Add CORS middleware if needed
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch, http.MethodOptions},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// Serve Swagger UI
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler(echoSwagger.URL("/swagger.yaml")))

	// Add a redirect for convenience
	e.GET("/swagger", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// Serve OpenAPI spec file directly
	e.GET("/swagger.yaml", func(c echo.Context) error {
		yamlFile, err := os.ReadFile("api/openapi.yaml")
		if err != nil {
			return c.String(http.StatusInternalServerError, "Could not read OpenAPI spec")
		}
		return c.Blob(http.StatusOK, "application/yaml", yamlFile)
	})

	// Register API handlers
	apigen.RegisterHandlersWithBaseURL(e, &server, "/api/v1")

	log.Info().Msgf("Starting server on port %d", cfg.Port)
	return e.Start(fmt.Sprintf(":%d", cfg.Port))
}
