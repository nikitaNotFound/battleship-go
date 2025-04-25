package main

import (
	"context"

	"github.com/nikitaNotFound/battleship/internal/webserver"
)

func main() {
	ctx := context.Background()
	webserver.Run(ctx)
}
