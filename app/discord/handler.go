package discord

import (
	"github.com/bwmarrin/discordgo"
)

func SlashCommandHandlerfunc(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}


	if !strings.HasPrefix(m.Content, "!") {
		return
	}

	args := strings.Split(m.Content[1:], " ")

	switch args[0] {
	case "ping":
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case "pong":
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	case "help":
		// s.ChannelMessageSend(m.ChannelID, "This is help")

		if len(args) > 1 {
			helpCommand(s, m, args[1])
		} else { // Help command without topic
			helpCommand(s, m, "")
		}
	default:
		s.ChannelMessageSend(m.ChannelID, errorMessage("Invalid command", "For a list of help topics, type !help"))
	}
}

func errorMessage(title string, message string) string {
	return "❌  **" + title + "**\n" + message
}
