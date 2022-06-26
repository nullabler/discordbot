package application

import "github.com/unixoff/discordbot/internal/handler"

func (app *App) addHandler() {
	app.handlers = append(app.handlers, handler.NewMessageHandler(app.config))
	app.handlers = append(app.handlers, handler.NewCommandHandler(app.config))
}
