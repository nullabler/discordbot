package discord

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func SlashCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, os.Getenv("TRIGGER")) {
		return
	}

	if !isAccess(m.Member.Roles) {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Error", "Permission denied"))
		return
	}

	args := strings.Split(m.Content[1:], " ")

	switch args[0] {
	case "ping":
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case "pong":
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	case "help":
		if len(args) > 1 {
			helpCommand(s, m, args[1])
		} else {
			helpCommand(s, m, "")
		}
	default:
		s.ChannelMessageSend(m.ChannelID, errorMessage("Invalid command", "For a list of help topics, type !help"))
	}
}

func MessageHandler(s * discordgo.Session, m *discordgo.MessageCreate) {
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
