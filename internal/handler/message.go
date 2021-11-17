package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/unixoff/discord-bot/context"
	"github.com/unixoff/discord-bot/discord"
)

type MessageHandler struct {
	ctx     *context.Context
	discord *discord.Discord
}

func NewMessageHandler(ctx *context.Context) *MessageHandler {
	return &MessageHandler{
		ctx:     ctx,
		discord: discord.New(ctx),
	}
}

func (handler *MessageHandler) Run(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !handler.discord.Init(s, m) {
		return
	}

	switch handler.discord.MessageContent() {
	case "привет", "всем привет", "хай", "hi", "hello":
		handler.discord.AddEmojiReaction("✌")
	case "спасибо", "спасибо за помощь", "спс":
		handler.discord.AddEmojiReaction("👍")
	}
}
