package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type State struct{
	session *discordgo.Session
	message *discordgo.MessageCreate
	channel *discordgo.Channel
	voiceConnect *discordgo.VoiceConnection
	args []string
}

func newState(s *discordgo.Session, m *discordgo.MessageCreate) *State {
	return &State{
		session: s,
		message: m,
	}
}

func (self *State) init() error {
	self.args = strings.Split(self.message.Content[1:], " ")
	return self.initChannel()
}


func (self *State) initChannel() (err error) {
	channel, err := self.session.Channel(self.message.ChannelID)
	if err != nil {
		return
	}

	self.channel = channel
	return
}

func (self *State) sessionGuilds() []*discordgo.Guild {
	return self.session.State.Guilds
}

func (self *State) messageAuth() *discordgo.User {
	return self.message.Author
}
