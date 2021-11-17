package handler

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/unixoff/discord-bot/config"
	"github.com/unixoff/discord-bot/context"
	"github.com/unixoff/discord-bot/discord"
)

type CommandHandler struct {
	ctx     *context.Context
	config  *config.Config
	discord *discord.Discord
}

func NewCommandHandler(ctx *context.Context) *CommandHandler {
	return &CommandHandler{
		ctx:     ctx,
		config:  ctx.Config(),
		discord: discord.New(ctx),
	}
}

func (handler *CommandHandler) Run(s *discordgo.Session, m *discordgo.MessageCreate) {

	if !handler.discord.Init(s, m) || !strings.HasPrefix(handler.discord.MessageContent(), handler.config.CommandTarget) {
		return
	}

	switch handler.discord.Args()[0] {
	case "ping", "pong":
		content := "Ping!"
		if handler.discord.Args()[0] == "ping" {
			content = "Pong!"
		}
		handler.discord.MessageSend(content)
	case "join":
		handler.discord.JoinToVoice()
	case "play":
		handler.discord.PlayMusic()
	case "pause":
		handler.discord.PauseMusic()
	case "resume":
		handler.discord.ResumeMusic()
	case "skip":
		handler.discord.SkipMusic()
	case "stop":
		handler.discord.StopMusic()
	case "disconnect":
		handler.discord.DisconnectMusic()
	default:
		handler.discord.InvalidCommand()
	}
}
