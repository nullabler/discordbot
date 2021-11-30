package handler

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/unixoff/discordbot/internal/config"
	"github.com/unixoff/discordbot/internal/context"
	"github.com/unixoff/discordbot/internal/discord"
	"github.com/unixoff/discordbot/internal/parse"
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
		content := "ping!"
		if handler.discord.Args()[0] == "ping" {
			content = "pong!"
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
	case "ver", "-v", "version":
		parse.Route(handler.discord)
	default:
		handler.discord.InvalidCommand()
	}
}
