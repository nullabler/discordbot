package handler

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/unixoff/discord-bot/config"
	"github.com/unixoff/discord-bot/context"
	"github.com/unixoff/discord-bot/service/discord"
)

type CommandHandler struct {
	ctx *context.Context
	config *config.Config
	discord *discord.Discord
}

func NewCommandHandler(ctx *context.Context) *CommandHandler {
	return &CommandHandler{
		ctx: ctx,
		config: ctx.Config(),
		discord: discord.New(),
	}
}

func (self *CommandHandler) Run(s *discordgo.Session, m *discordgo.MessageCreate) {

	if !self.discord.Init(s, m) || !strings.HasPrefix(self.discord.MessageContent(), self.config.CommandTarget) {
		return
	}

	switch self.discord.Args()[0] {
	case "ping", "pong":
		content := "Ping!"
		if self.discord.Args()[0] == "ping" {
			content = "Pong!"
		}
		s.ChannelMessageSend(m.ChannelID, content)
	// case "play", "stop", "skip", "pause", "disconnect", "join", "radio":
	// 	if !isAccessForMusic() {
	// 		permissionDeniedMessage()
	// 		return
	// 	}
	// 	music.Handler(s, m, args)
	// case "help":
	// 	if len(args) > 1 {
	// 		helpCommand(s, m, args[1])
	// 	} else {
	// 		helpCommand(s, m, "")
	// 	}
	// default:
	// 	s.ChannelMessageSend(m.ChannelID, errorMessage("Invalid command", "For a list of help topics, type !help"))
	// }
	}
}
