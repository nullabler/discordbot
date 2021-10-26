package application

import (
	commandHandler "github.com/unixoff/discord-bot/handlers/command_handler"
	messageHandler "github.com/unixoff/discord-bot/handlers/message_handler"
)

func (app *App) addService() {
	app.handlers = append(app.handlers, messageHandler.New(app.ctx))
	app.handlers = append(app.handlers, commandHandler.New(app.ctx))
}
