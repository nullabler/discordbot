package music

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	// "log"
	// _"strconv"
	// "strings"
)

func setVar(s *discordgo.Session, m *discordgo.MessageCreate) {
	session, message = s, m
}

func PlayCommand(s *discordgo.Session, m *discordgo.MessageCreate, query string) {
	setVar(s, m)
	guildID := SearchGuild(m.ChannelID)
	// voice := voiceInstances[guildID]

	mutex.Lock()
	v := new(VoiceInstance)
	voiceInstances[guildID] = v
	v.guildID = guildID
	v.session = s
	mutex.Unlock()

	voiceChannelID := SearchVoiceChannel(guildID, m.Author.ID)
	v.voice ,_ = session.ChannelVoiceJoin(guildID, voiceChannelID, false, false)

	ChMessageSend(m.ChannelID, "[**Music**] Join")


	song := Song{
		m.ChannelID,
		"De Rien - Night Bus",
		m.Author.ID,
		"vV9MvDz5aaU",
		"De Rien - Night Bus",
		"3:13",
		// "https://youtu.be/vV9MvDz5aaU",
		"https://r7---sn-gvnuxaxjvh-nufl.googlevideo.com/videoplayback?expire=1634498166&ei=FiJsYa7gMZPT7QT9tYPIAQ&ip=45.137.189.179&id=o-AMmwA2MqU-1ke7-fS7ShrSDfOEnERqW3pMSO0Q_7Sk_V&itag=18&source=youtube&requiressl=yes&vprv=1&mime=video%2Fmp4&ns=02kzW4XsDyQHHu4RX8sBCswG&gir=yes&clen=3384535&ratebypass=yes&dur=193.399&lmt=1634030685468752&fexp=24001373%2C24007246&c=WEB&txp=6310222&n=wxUj3n6EZjHZP5y-1&sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Cvprv%2Cmime%2Cns%2Cgir%2Cclen%2Cratebypass%2Cdur%2Clmt&sig=AOq0QJ8wRQIgW1sEpeClYBRgd7-tJeGPyCoWo7GP8GAfMEzQiuwchMMCIQDKjnpfe-nCNL4EQQDX25nQQ5eJGwGBcO3lMGFh-MpudA==&title=De_Rien_-_Night_Bus&cms_redirect=yes&mh=3K&mip=178.34.163.128&mm=31&mn=sn-gvnuxaxjvh-nufl&ms=au&mt=1634476137&mv=m&mvi=7&pcm2cms=yes&pl=24&lsparams=mh,mip,mm,mn,ms,mv,mvi,pcm2cms,pl&lsig=AG3C_xAwRQIhAItSm-PLL_2iLdYyIWKHgg1YWXJDb36zEo_nWMPaevUZAiAbmHzIIE4T7QWNKJ83agrWcHXD_I90qkGG48Fs1DXGmQ%3D%3D",
	}

	song_struct := PkgSong{
		data: song,
		v: v,
	}

	fmt.Println(song_struct)
	go func() {
		songSignal <- song_struct
	}()
}

func DisconnectCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	guildID := SearchGuild(m.ChannelID)
	v := voiceInstances[guildID]
	v.voice.Disconnect()
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
