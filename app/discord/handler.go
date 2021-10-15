package discord

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	RoleList []string
)

func SlashCommandHandlerfunc(s *discordgo.Session, m *discordgo.MessageCreate) {
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

func isAccess(memberRoleList []string) bool {
	if len(RoleList) == 0 {
		RoleList =  strings.Split(os.Getenv("ROLE_LIST"), ":")
	}

	for _, memberRole := range memberRoleList {
		for _, role := range RoleList {
			if memberRole == role {
				return true
			}
		}
	}
	return false
}
