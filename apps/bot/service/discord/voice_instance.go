package discord

import (
	"log"
	"os/exec"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
)

type VoiceInstance struct {
	state *State

	// voice      *discordgo.VoiceConnection
	session    *discordgo.Session
	encoder    *dca.EncodeSession
	stream     *dca.StreamingSession
	run        *exec.Cmd
	queueMutex sync.Mutex
	audioMutex sync.Mutex
	nowPlaying Song
	queue      []Song
	recv       []int16
	guildID    string
	channelID  string
	speaking   bool
	pause      bool
	stop       bool
	skip       bool
	radioFlag  bool
}

func newVoiceInstance(state *State) *VoiceInstance {
	return &VoiceInstance{
		state: state,
	}
}

func (self *VoiceInstance) JoinToVoice() error {
	voiceChannelID := self.searchVoiceChannelID()
	voiceConnect, err := self.state.session.ChannelVoiceJoin(self.state.channel.GuildID, voiceChannelID, false, false)
	if err != nil {
		log.Println("Error: channel voice join", err)
		return err
	}
	self.state.voiceConnect = voiceConnect
	return nil
}

func (self *VoiceInstance) searchVoiceChannelID() (voiceChannelID string) {
	for _, guild := range self.state.sessionGuilds() {
		if guild.ID != self.state.channel.GuildID {
			continue
		}

		for _, voiceState := range guild.VoiceStates {
			if voiceState.UserID == self.state.messageAuth().ID {
				return voiceState.ChannelID
			}
		}
	}

	return ""
}
