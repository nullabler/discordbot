package handler

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/unixoff/discordbot/internal/config"
	"github.com/unixoff/discordbot/internal/discord"
	"github.com/unixoff/discordbot/internal/parse"
)

type CommandHandler struct {
	config  *config.Config
	discord *discord.Discord
}

func NewCommandHandler(conf *config.Config) *CommandHandler {
	return &CommandHandler{
		config:  conf,
		discord: discord.New(conf),
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
	case "ver", "-v", "version":
		parse.Route(handler.discord)
	default:
		handler.discord.InvalidCommand()
	}
}
