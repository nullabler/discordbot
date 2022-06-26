package application

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/unixoff/discordbot/internal/config"
	"github.com/unixoff/discordbot/internal/handler"
)

type App struct {
	config *config.Config
	handlers []handler.HandlerInterface
	signal chan os.Signal
}

func New() *App {
	app := &App{
		config: config.New(),
		signal: make(chan os.Signal, 1),
	}

	app.addHandler()

	return app
}

func (app *App) Run() {
	signal.Notify(app.signal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	discordSession, err := discordgo.New("Bot " + app.config.DiscordToken)
	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	for _, handler := range app.handlers {
		discordSession.AddHandler(handler.Run)
	}

	err = discordSession.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return
	}

	_, err = discordSession.User("@me")
	if err != nil {
		log.Println("FATA:", err)
		return
	}

	log.Println("Bot is now running.  Press CTRL-C to exit.")

	<-app.signal
	discordSession.Close()
}
