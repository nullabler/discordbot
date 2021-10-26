package main

import (
	"github.com/unixoff/discord-bot/application"
	"github.com/unixoff/discord-bot/context"
)


func main() {
	ctx := context.New()
	app := application.New(ctx)
	app.Run()
}
