package discord

import (
	"log"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type Discord struct{
	state *State
	voiceInstanceList map[string]*VoiceInstance
	mutex sync.Mutex
	currentVoiceInstance *VoiceInstance
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

	if vInstance, ok := self.voiceInstanceList[self.state.channel.GuildID]; ok {
		self.currentVoiceInstance = vInstance
	}

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

func (self *Discord) MessageSend(content string) {
	self.state.session.ChannelMessageSend(self.state.message.ChannelID, content)
}

func (self *Discord) JoinToVoice() {
	if self.currentVoiceInstance == nil {
		self.mutex.Lock()
		self.currentVoiceInstance = newVoiceInstance(self.state)
		self.voiceInstanceList[self.state.channel.GuildID] = self.currentVoiceInstance
		self.mutex.Unlock()
	}

	self.currentVoiceInstance.JoinToVoice()
}
