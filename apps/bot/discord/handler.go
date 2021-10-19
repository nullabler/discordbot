package discord

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/unixoff/discord-bot/music"
)

func setVar(s *discordgo.Session, m *discordgo.MessageCreate) {
	session, message = s, m

}

func SlashCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	setVar(s, m)
	if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, os.Getenv("TRIGGER")) || !isAccess() {
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
	case "play", "stop", "skip", "pause", "disconnect", "join", "radio":
		if !isAccessForMusic() {
			permissionDeniedMessage()
			return
		}
		music.Handler(s, m, args)
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

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
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
