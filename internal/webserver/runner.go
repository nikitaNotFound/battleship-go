package webserver

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
)

func Run(ctx context.Context) error {
	cfg, err := GetWebServerConfig()
	if err != nil {
		return err
	}

	e := echo.New()

	

	return e.Start(fmt.Sprintf(":%d", cfg.Port))
}
