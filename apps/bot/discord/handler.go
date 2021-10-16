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

	if !isAccess() {
		permissionDeniedMessage()
		return
	}

	args := strings.Split(m.Content[1:], " ")

	switch args[0] {
	case "ping":
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case "pong":
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	case "play":
		if !isAccessForMusic(os.Getenv("MUSIC_CHANNEL")) || len(args) != 2 {
			permissionDeniedMessage()
			return
		}
		music.PlayCommand(s, m, args[1])
	case "disconnect":
		if !isAccessForMusic(os.Getenv("MUSIC_CHANNEL")) {
			permissionDeniedMessage()
			return
		}
		music.DisconnectCommand(s, m)
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
