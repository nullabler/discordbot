package music

import (
	"github.com/bwmarrin/discordgo"
	"sync"
)

var (
	session        *discordgo.Session
	message		   *discordgo.MessageCreate
	// voice		   *discordgo.VoiceConnection

	voiceInstances = map[string]*VoiceInstance{}
	purgeTime      int64
	purgeQueue     []PurgeMessage
	mutex          sync.Mutex
	songSignal     chan PkgSong
	radioSignal    chan PkgRadio
	//ignore            = map[string]bool{}
)
