package handlers

import "github.com/bwmarrin/discordgo"

type HandlerInterface interface {
	Run(s *discordgo.Session, m *discordgo.MessageCreate)
}
