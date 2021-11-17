package discord

import (
	"io"
	"log"
	"os/exec"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
)

type VoiceInstance struct {
	state *State

	connect    *discordgo.VoiceConnection
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

func (v *VoiceInstance) join() error {
	voiceChannelID := v.searchVoiceChannelID()
	voiceConnect, err := v.state.session.ChannelVoiceJoin(v.state.channel.GuildID, voiceChannelID, false, false)
	if err != nil {
		log.Println("Error: channel voice join", err)
		return err
	}
	v.connect = voiceConnect
	return nil
}

func (v *VoiceInstance) searchVoiceChannelID() (voiceChannelID string) {
	for _, guild := range v.state.sessionGuilds() {
		if guild.ID != v.state.channel.GuildID {
			continue
		}

		for _, voiceState := range guild.VoiceStates {
			if voiceState.UserID == v.state.messageAuth().ID {
				return voiceState.ChannelID
			}
		}
	}

	return ""
}

func (v *VoiceInstance) Stop() {
	v.stop = true
	if v.encoder != nil {
		v.encoder.Cleanup()
	}
}

func (v *VoiceInstance) Skip() bool {
	if v.speaking {
		if v.pause {
			return true
		} else {
			if v.encoder != nil {
				v.encoder.Cleanup()
			}
		}
	}
	return false
}

// Pause pause the audio
func (v *VoiceInstance) Pause() {
	v.pause = true
	if v.stream != nil {
		v.stream.SetPaused(true)
	}
}

// Resume resume the audio
func (v *VoiceInstance) Resume() {
	v.pause = false
	if v.stream != nil {
		v.stream.SetPaused(false)
	}
}

func (v *VoiceInstance) PlayQueue(song Song) {
	// add song to queue
	v.QueueAdd(song)
	if v.speaking {
		// the bot is playing
		return
	}
	go func() {
		v.audioMutex.Lock()
		defer v.audioMutex.Unlock()
		for {
			if len(v.queue) == 0 {
				// dg.UpdateStatus(0, o.DiscordStatus)
				//ChMessageSend(v.nowPlaying.ChannelID, "[**Music**] End of queue!")
				return
			}
			v.nowPlaying = v.QueueGetSong()
			//go ChMessageSend(v.nowPlaying.ChannelID, "[**Music**] Playing, **`"+
			//v.nowPlayig.Title+"`  -  `("+v.nowPlaying.Duration+")`  -  **<@"+v.nowPlaying.ID+">\n") /[>`"+ v.nowPlaying.User +"`***")
			// If monoserver
			//if o.DiscordPlayStatus {
			// dg.UpdateStatus(0, v.nowPlaying.Title)
			//}
			v.stop = false
			v.skip = false
			v.speaking = true
			v.pause = false
			v.connect.Speaking(true)

			v.DCA(v.nowPlaying.VideoURL)

			v.QueueRemoveFisrt()
			if v.stop {
				v.QueueRemove()
			}

			v.stop = false
			v.skip = false
			v.speaking = false
			v.connect.Speaking(false)
		}
	}()
}

func (v *VoiceInstance) DCA(url string) {
	opts := dca.StdEncodeOptions
	opts.RawOutput = true
	opts.Bitrate = 96
	opts.Application = "lowdelay"
	opts.BufferedFrames = 100

	encodeSession, err := dca.EncodeFile(url, opts)
	if err != nil {
		log.Println("FATA: Failed creating an encoding session: ", err)
	}
	v.encoder = encodeSession
	done := make(chan error)

	stream := dca.NewStream(encodeSession, v.connect, done)
	v.stream = stream
	for {
		select {
		case err := <-done:
			if err != nil && err != io.EOF {
				log.Println("FATA: An error occured", err)
			}
			// Clean up incase something happened and ffmpeg is still running
			encodeSession.Cleanup()
			return
		}
	}
}

func (v *VoiceInstance) QueueGetSong() (song Song) {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	if len(v.queue) != 0 {
		return v.queue[0]
	}
	return
}

// QueueAdd
func (v *VoiceInstance) QueueAdd(song Song) {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	v.queue = append(v.queue, song)
}

// QueueRemoveFirst
func (v *VoiceInstance) QueueRemoveFisrt() {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	if len(v.queue) != 0 {
		v.queue = v.queue[1:]
	}
}

// QueueRemoveIndex
func (v *VoiceInstance) QueueRemoveIndex(k int) {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	if len(v.queue) != 0 && k <= len(v.queue) {
		v.queue = append(v.queue[:k], v.queue[k+1:]...)
	}
}

// QueueRemoveUser
func (v *VoiceInstance) QueueRemoveUser(user string) {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	queue := v.queue
	v.queue = []Song{}
	if len(v.queue) != 0 {
		for _, q := range queue {
			if q.User != user {
				v.queue = append(v.queue, q)
			}
		}
	}
}

// QueueRemoveLast
func (v *VoiceInstance) QueueRemoveLast() {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	if len(v.queue) != 0 {
		v.queue = append(v.queue[:len(v.queue)-1], v.queue[len(v.queue):]...)
	}
}

// QueueClean
func (v *VoiceInstance) QueueClean() {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	// hold the actual song in the queue
	v.queue = v.queue[:1]
}

// QueueRemove
func (v *VoiceInstance) QueueRemove() {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	v.queue = []Song{}
}
