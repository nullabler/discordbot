package discord

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/unixoff/discordbot/internal/config"
	"github.com/unixoff/discordbot/internal/context"
)

type Discord struct {
	indexInvalidCommand  int
	config               *config.Config
	state                *State
	voiceInstanceList    map[string]*VoiceInstance
	mutex                sync.Mutex
	currentVoiceInstance *VoiceInstance
	youtube              *Youtube
	songSignal           chan PkgSong
}

func New(ctx *context.Context) *Discord {
	return &Discord{
		config:              ctx.Config(),
		voiceInstanceList:   make(map[string]*VoiceInstance),
		youtube:             newYoutube(),
		songSignal:          make(chan PkgSong),
		indexInvalidCommand: 0,
	}
}

func (discord *Discord) Init(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if m.Author.Bot || len(m.Attachments) > 0 {
		return false
	}

	discord.state = newState(s, m)
	discord.youtube.init(discord.state)
	if err := discord.state.init(); err != nil {
		log.Println(err)
		return false
	}

	if vInstance, ok := discord.voiceInstanceList[discord.state.channel.GuildID]; ok {
		discord.currentVoiceInstance = vInstance
	}

	go discord.globalPlayMusic()

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

func (discord *Discord) JoinToVoice() {
	if !discord.isAccessForMusic() {
		return
	}

	if discord.currentVoiceInstance == nil {
		discord.mutex.Lock()
		discord.currentVoiceInstance = newVoiceInstance(discord.state)
		discord.voiceInstanceList[discord.state.channel.GuildID] = discord.currentVoiceInstance
		discord.mutex.Unlock()
	}

	discord.currentVoiceInstance.join()
}

func (discord *Discord) PlayMusic() {
	if !discord.isAccessForMusic() {
		return
	}

	if discord.currentVoiceInstance == nil {
		discord.JoinToVoice()
	} else {
		if err := discord.currentVoiceInstance.join(); err != nil {
			return
		}
	}

	pkgSong, err := discord.youtube.find(discord.Args()[1])
	pkgSong.v = discord.currentVoiceInstance
	if err != nil {
		discord.MessageSend("Youtube error")
		return
	}

	go func() {
		discord.songSignal <- pkgSong
	}()
}

func (discord *Discord) PauseMusic() {
	if !discord.isAccessForMusic() {
		return
	}

	if discord.currentVoiceInstance == nil {
		log.Println("INFO: The bot is not joined in a voice channel")
		return
	}

	if !discord.currentVoiceInstance.speaking {
		discord.MessageSend("I'm not playing nothing!")
		return
	}

	if !discord.currentVoiceInstance.pause {
		discord.currentVoiceInstance.Pause()
		discord.MessageSend("I'm `PAUSED` now!")
	}
}

func (discord *Discord) ResumeMusic() {
	if !discord.isAccessForMusic() {
		return
	}

	if discord.currentVoiceInstance == nil {
		log.Println("The bot is not joined in a voice channel")
		return
	}

	if !discord.currentVoiceInstance.speaking {
		discord.MessageSend("I'm not playing nothing!")
		return
	}

	if discord.currentVoiceInstance.pause {
		discord.currentVoiceInstance.Resume()
		discord.MessageSend("I'm `RESUMED` now!")
	}
}

func (discord *Discord) SkipMusic() {
	if !discord.isAccessForMusic() {
		return
	}

	if discord.currentVoiceInstance == nil {
		log.Println("INFO: The bot is not joined in a voice channel")
		discord.MessageSend("I need join in a voice channel!")
		return
	}

	if len(discord.currentVoiceInstance.queue) == 0 {
		log.Println("INFO: The queue is empty.")
		discord.MessageSend("Currently there's no music playing, add some? ;)")
		return
	}

	if discord.currentVoiceInstance.Skip() {
		discord.MessageSend("I'm `PAUSED`, please `play` first.")
	}
}

func (discord *Discord) StopMusic() {
	if !discord.isAccessForMusic() || discord.currentVoiceInstance == nil {
		return
	}

	if discord.currentVoiceInstance.connect.ChannelID != discord.state.channel.ID {
		discord.MessageSend("<@" + discord.state.messageAuth().ID + "> You need to join in my voice channel for send stop!")
		return
	}
	discord.currentVoiceInstance.Stop()
	log.Println("INFO: The bot stop play audio")
	discord.MessageSend("I'm stoped now!")
}

func (discord *Discord) DisconnectMusic() {
	if !discord.isAccessForMusic() || discord.currentVoiceInstance == nil {
		return
	}

	if discord.currentVoiceInstance.connect.Ready {
		discord.currentVoiceInstance.connect.Disconnect()
	}
	discord.currentVoiceInstance.Stop()
	time.Sleep(200 * time.Millisecond)
	discord.mutex.Lock()
	delete(discord.voiceInstanceList, discord.state.channel.GuildID)
	discord.mutex.Unlock()
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

func (discord *Discord) globalPlayMusic() {
	for {
		select {
		case song := <-discord.songSignal:
			if song.v.radioFlag {
				song.v.Stop()
				time.Sleep(200 * time.Millisecond)
			}

			go song.v.PlayQueue(song.data)
		}
	}
}

func (discord *Discord) isAccessForMusic() bool {
	for _, channelID := range discord.config.MusicChannelList {
		if channelID == discord.state.message.ChannelID {
			return true
		}
	}

	discord.AddEmojiReaction("🚫")
	return false
}
