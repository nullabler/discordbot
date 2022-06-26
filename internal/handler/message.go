package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/unixoff/discordbot/internal/config"
	"github.com/unixoff/discordbot/internal/discord"
)

type MessageHandler struct {
	conf *config.Config
	discord *discord.Discord
}

func NewMessageHandler(conf *config.Config) *MessageHandler {
	return &MessageHandler{
		conf:     conf,
		discord: discord.New(conf),
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
