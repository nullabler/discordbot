package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/unixoff/discord-bot/config"
	"github.com/unixoff/discord-bot/context"
)

type CommandHandler struct {
	ctx *context.Context
	config *config.Config
}

func New(ctx *context.Context) *CommandHandler {
	return &CommandHandler{
		ctx: ctx,
		config: ctx.Config(),
	}
}

func (self *CommandHandler) Run(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, self.config.CommandTarget) {
		return
	}

	args := strings.Split(m.Content[1:], " ")
	switch args[0] {
	case "ping", "pong":
		content := "Ping!"
		if args[0] == "ping" {
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
