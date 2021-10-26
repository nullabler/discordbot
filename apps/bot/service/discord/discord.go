package discord

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Discord struct{
	state *State
	voiceInstanceList map[string]*VoiceInstance
}

func New() *Discord {
	return &Discord{}
}

func (self *Discord) Init(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	self.state = newState(s, m)
	if err := self.state.init(); err != nil {
		log.Println(err)
		return false
	}
	self.state.voiceInstance = self.voiceInstanceList[self.state.channel.GuildID]
	return m.Author.ID != s.State.User.ID
}

func (self *Discord) MessageContent() string {
	return strings.ToLower(self.state.message.Content)
}

func (self *Discord) AddEmojiReaction(emojiID string) {
	self.state.session.MessageReactionAdd(self.state.message.ChannelID, self.state.message.ID, emojiID)
}

func (self *Discord) Args() []string {
	return self.state.args
}
