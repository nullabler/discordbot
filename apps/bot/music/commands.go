package music

import (
	"time"

	"github.com/bwmarrin/discordgo"
	// "log"
	// _"strconv"
	// "strings"
)

func setVar(s *discordgo.Session, m *discordgo.MessageCreate) {
	session, message = s, m
}

func JoinCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	setVar(s, m)
	guildID := SearchGuild(m.ChannelID)

	vInstance := voiceInstances[guildID]
	if vInstance == nil {
		mutex.Lock()
		vInstance = new(VoiceInstance)
		voiceInstances[guildID] = vInstance
		vInstance.guildID = guildID
		vInstance.session = s
		mutex.Unlock()
	}
	ChVoiceJoin(guildID, vInstance)
}

func PlayCommand(s *discordgo.Session, m *discordgo.MessageCreate, query string) {
	setVar(s, m)
	guildID := SearchGuild(m.ChannelID)
	vInstance := voiceInstances[guildID]


	if vInstance == nil {
		JoinCommand(s, m)
	} else {
		if err := ChVoiceJoin(guildID, vInstance); err != nil {
			return
		}
	}

	pkgSong, err := youtubeFind(query, vInstance)
	if err != nil {
		ChMessageSend(m.ChannelID, "[**Music**] youtube error")
		return
	}

	go func() {
		songSignal <- pkgSong
	}()
}

func DisconnectCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	guildID := SearchGuild(m.ChannelID)
	vInstance := voiceInstances[guildID]
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
// func PlayReporter(v *VoiceInstance, m *discordgo.MessageCreate) {
// 	log.Println("INFO:", m.Author.Username, "send 'play'")
// 	if v == nil {
// 		log.Println("INFO: The bot is not joined in a voice channel")
// 		ChMessageSend(m.ChannelID, "[**Music**] I need join in a voice channel!")
// 		return
// 	}
// 	if len(strings.Fields(m.Content)) < 2 {
// 		ChMessageSend(m.ChannelID, "[**Music**] You need specify a name or URL.")
// 		return
// 	}
// 	// if the user is not a voice channel not accept the command
// 	voiceChannelID := SearchVoiceChannel(m.Author.ID)
// 	if v.voice.ChannelID != voiceChannelID {
// 		ChMessageSend(m.ChannelID, "[**Music**] <@"+m.Author.ID+"> You need to join in my voice channel for send play!")
// 		return
// 	}
// 	// send play my_song_youtube
// 	command := strings.SplitAfter(m.Content, strings.Fields(m.Content)[0])
// 	query := strings.TrimSpace(command[1])
// 	// song, err := YoutubeFind(query, v, m)
// 	// if err != nil || song.data.ID == "" {
// 	// 	log.Println("ERROR: Youtube search: ", err)
// 	// 	ChMessageSend(m.ChannelID, "[**Music**] I can't found this song!")
// 	// 	return
// 	// }

// 	ChMessageSend(m.ChannelID, "[**Music**] **`User`** has added , **`Title`** to the queue. **`(Duration)` `[strconv.Itoa(len(v.queue))]`**"+query)
// 	//***`"+ song.data.User +"`***
// 	// ChMessageSend(m.ChannelID, "[**Music**] **`"+song.data.User+"`** has added , **`"+
// 	// 	song.data.Title+"`** to the queue. **`("+song.data.Duration+")` `["+strconv.Itoa(len(v.queue))+"]`**")
// 	// go func() {
// 	// 	songSignal <- song
// 	// }()
// }
