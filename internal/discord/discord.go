package discord

import (
	"log"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/unixoff/discordbot/internal/config"
)

type Discord struct {
	indexInvalidCommand  int
	config               *config.Config
	state                *State
	mutex                sync.Mutex
}

func New(conf *config.Config) *Discord {
	return &Discord{
		config:              conf,
		indexInvalidCommand: 0,
	}
}

func (discord *Discord) Init(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if m.Author.Bot || len(m.Attachments) > 0 {
		return false
	}

	discord.state = newState(s, m)
	if err := discord.state.init(); err != nil {
		log.Println(err)
		return false
	}

	return m.Author.ID != s.State.User.ID
}

func (discord *Discord) MessageContent() string {
	return strings.ToLower(discord.state.message.Content)
}

func (discord *Discord) AddEmojiReaction(emojiID string) {
	discord.state.session.MessageReactionAdd(discord.state.message.ChannelID, discord.state.message.ID, emojiID)
}

func (discord *Discord) Args() []string {
	return discord.state.args
}

func (discord *Discord) MessageSend(content string) {
	discord.state.session.ChannelMessageSend(discord.state.message.ChannelID, content)
}

func (discord *Discord) HasChannelID(channelID string) bool {
	return discord.state.message.ChannelID == channelID
}

func (discord *Discord) InvalidCommand() {
	switch discord.indexInvalidCommand {
	case 0:
		discord.AddEmojiReaction("🙉")
	case 1:
		discord.AddEmojiReaction("🙈")
	default:
		discord.AddEmojiReaction("🙊")
	}
	discord.indexInvalidCommand++
	if discord.indexInvalidCommand > 2 {
		discord.indexInvalidCommand = 0
	}
}

//func (discord *Discord) globalPlayMusic() {
	//for {
		//select {
		//case song := <-discord.songSignal:
			//if song.v.radioFlag {
				//song.v.Stop()
				//time.Sleep(200 * time.Millisecond)
			//}

			//go song.v.PlayQueue(song.data)
		//}
	//}
//}

func (discord *Discord) isAccessForMusic() bool {
	for _, channelID := range discord.config.MusicChannelList {
		if channelID == discord.state.message.ChannelID {
			return true
		}
	}

	discord.AddEmojiReaction("🚫")
	return false
}
