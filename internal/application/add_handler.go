package application

import "github.com/unixoff/discord-bot/handler"

func (app *App) addHandler() {
	app.handlers = append(app.handlers, handler.NewMessageHandler(app.ctx))
	app.handlers = append(app.handlers, handler.NewCommandHandler(app.ctx))
}
