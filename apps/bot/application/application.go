package application

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/unixoff/discord-bot/config"
	"github.com/unixoff/discord-bot/context"
	"github.com/unixoff/discord-bot/handler"
)

type App struct {
	ctx               *context.Context
	config            *config.Config
	// voiceInstanceList map[string]*VoiceInstance

	handlers []handler.HandlerInterface

	sigs chan os.Signal
	Quit chan bool
}

func New(ctx *context.Context) *App {
	app := &App{
		ctx: ctx,
		sigs: make(chan os.Signal, 1),
		Quit: make(chan bool, 1),
	}

	app.init()

	return app
}

func (app *App) init() {
	app.config = config.New()
	app.ctx.SetConfig(app.config)

	app.addHandler()
}


func (app *App) Run() {
	signal.Notify(app.sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	go func() {
		<-app.sigs
		app.Quit <- true
	}()

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

	<-app.Quit
	discordSession.Close()
}
