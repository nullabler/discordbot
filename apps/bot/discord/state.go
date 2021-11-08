package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type State struct {
	session *discordgo.Session
	message *discordgo.MessageCreate
	channel *discordgo.Channel
	args    []string
}

func newState(s *discordgo.Session, m *discordgo.MessageCreate) *State {
	return &State{
		session: s,
		message: m,
	}
}

func (state *State) init() error {
	state.args = strings.Split(state.message.Content[1:], " ")
	return state.initChannel()
}

func (state *State) initChannel() (err error) {
	channel, err := state.session.Channel(state.message.ChannelID)
	if err != nil {
		return
	}

	state.channel = channel
	return
}

func (state *State) sessionGuilds() []*discordgo.Guild {
	return state.session.State.Guilds
}

func (state *State) messageAuth() *discordgo.User {
	return state.message.Author
}
