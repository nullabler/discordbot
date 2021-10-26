package message

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/unixoff/discord-bot/config"
	"github.com/unixoff/discord-bot/context"
)

type MessageHandler struct {
	ctx *context.Context
	config *config.Config
}

func New(ctx *context.Context) *MessageHandler {
	return &MessageHandler{
		ctx: ctx,
		config: ctx.Config(),
	}
}

func (self *MessageHandler) Run(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch strings.ToLower(m.Content) {
	case "привет", "всем привет", "хай", "hi", "hello":
		s.MessageReactionAdd(m.ChannelID, m.ID, "✌")
	case "спасибо", "спасибо за помощь", "спс":
		s.MessageReactionAdd(m.ChannelID, m.ID, "👍")
	}
}
