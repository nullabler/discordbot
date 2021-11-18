package main

import (
	"github.com/unixoff/discordbot/internal/application"
	"github.com/unixoff/discordbot/internal/context"
)

func main() {
	ctx := context.New()
	app := application.New(ctx)
	app.Run()
}
