package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/unixoff/discord-bot/context"
	"github.com/unixoff/discord-bot/service/discord"
)

type MessageHandler struct {
	ctx *context.Context
	discord *discord.Discord
}

func NewMessageHandler(ctx *context.Context) *MessageHandler {
	return &MessageHandler{
		ctx: ctx,
		discord: discord.New(),
	}
}

func (self *MessageHandler) Run(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !self.discord.Init(s, m) {
		return
	}

	switch self.discord.MessageContent() {
	case "привет", "всем привет", "хай", "hi", "hello":
		self.discord.AddEmojiReaction("✌")
	case "спасибо", "спасибо за помощь", "спс":
		self.discord.AddEmojiReaction("👍")
	}
}
