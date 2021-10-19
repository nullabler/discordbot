package music

import (
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func setVar(s *discordgo.Session, m *discordgo.MessageCreate) {
	session, message = s, m
}

func JoinCommand(guildID string) {
	vInstance := voiceInstances[guildID]
	if vInstance == nil {
		mutex.Lock()
		vInstance = new(VoiceInstance)
		voiceInstances[guildID] = vInstance
		vInstance.guildID = guildID
		vInstance.session = session
		mutex.Unlock()
	}
	ChVoiceJoin(guildID, vInstance)
}

func PlayCommand(guildID string, vInstance *VoiceInstance, query string) {
	if vInstance == nil {
		JoinCommand(guildID)
		vInstance = voiceInstances[guildID]
	} else {
		if err := ChVoiceJoin(guildID, vInstance); err != nil {
			return
		}
	}

	pkgSong, err := youtubeFind(query, vInstance)
	if err != nil {
		ChMessageSend(message.ChannelID, "[**Music**] youtube error")
		return
	}

	go func() {
		songSignal <- pkgSong
	}()
}

func PauseCommand(vInstance *VoiceInstance) {
	log.Println("INFO:", message.Author.Username, "send 'pause'")
	if vInstance == nil {
		log.Println("INFO: The bot is not joined in a voice channel")
		return
	}
	if !vInstance.speaking {
		ChMessageSend(message.ChannelID, "[**Music**] I'm not playing nothing!")
		return
	}
	if !vInstance.pause {
		vInstance.Pause()
		ChMessageSend(message.ChannelID, "[**Music**] I'm `PAUSED` now!")
	}
}

func ResumeCommand(vInstance *VoiceInstance) {
	log.Println("INFO:", message.Author.Username, "send 'resume'")
	if vInstance == nil {
		log.Println("INFO: The bot is not joined in a voice channel")
		ChMessageSend(message.ChannelID, "[**Music**] I need join in a voice channel!")
		return
	}
	if !vInstance.speaking {
		ChMessageSend(message.ChannelID, "[**Music**] I'm not playing nothing!")
		return
	}
	if vInstance.pause {
		vInstance.Resume()
		ChMessageSend(message.ChannelID, "[**Music**] I'm `RESUMED` now!")
	}
}

func SkipCommand(vInstance *VoiceInstance) {
	log.Println("INFO:", message.Author.Username, "send 'skip'")
	if vInstance == nil {
		log.Println("INFO: The bot is not joined in a voice channel")
		ChMessageSend(message.ChannelID, "[**Music**] I need join in a voice channel!")
		return
	}
	if len(vInstance.queue) == 0 {
		log.Println("INFO: The queue is empty.")
		ChMessageSend(message.ChannelID, "[**Music**] Currently there's no music playing, add some? ;)")
		return
	}
	if vInstance.Skip() {
		ChMessageSend(message.ChannelID, "[**Music**] I'm `PAUSED`, please `play` first.")
	}
}

func StopCommand(guildID string, vInstance *VoiceInstance) {
	if vInstance == nil {
		return
	}

	voiceChannelID := SearchVoiceChannel(guildID, message.Author.ID)
	if vInstance.voice.ChannelID != voiceChannelID {
		ChMessageSend(message.ChannelID, "[**Music**] <@"+message.Author.ID+"> You need to join in my voice channel for send stop!")
		return
	}
	vInstance.Stop()
	log.Println("INFO: The bot stop play audio")
	ChMessageSend(message.ChannelID, "[**Music**] I'm stoped now!")
}


func DisconnectCommand(guildID string, vInstance *VoiceInstance) {
	if vInstance == nil {
		return
	}

	if vInstance.voice.Ready {
		vInstance.voice.Disconnect()
	}
	vInstance.Stop()
	time.Sleep(200 * time.Millisecond)
	mutex.Lock()
	delete(voiceInstances, guildID)
	mutex.Unlock()
}

func RadioCommand(vInstance *VoiceInstance) {
	log.Println("INFO:", message.Author.Username, "send 'radio'")
	if vInstance == nil {
		log.Println("INFO: The bot is not joined in a voice channel")
		ChMessageSend(message.ChannelID, "[**Music**] I need join in a voice channel!")
		return
	}
	if len(strings.Fields(message.Content)) < 2 {
		ChMessageSend(message.ChannelID, "[**Music**] You need to specify a url!")
		return
	}
	radio := PkgRadio{"", vInstance}
	radio.data = strings.Fields(message.Content)[1]

	go func() {
		radioSignal <- radio
	}()
	ChMessageSend(message.ChannelID, "[**Music**] **`"+message.Author.Username+"`** I'm playing a radio now!")
}
