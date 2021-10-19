package music

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

var o = &Options{}

func InitRoutine() {
	songSignal = make(chan PkgSong)
	radioSignal = make(chan PkgRadio)
	go GlobalPlay(songSignal)
	go GlobalRadio(radioSignal)
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	setVar(s, m)
	guildID, err := SearchGuild()
	if err != nil {
		log.Println(err)
		return
	}
	vInstance := voiceInstances[guildID]

	switch args[0] {
	case "join":
		JoinCommand(guildID)
	case "play":
		if len(args) == 1 {
			ResumeCommand(vInstance)
			return
		}

		if len(args) != 2 {
			ChMessageSend(m.ChannelID, "Not found URL")
			return
		}
		PlayCommand(guildID, vInstance, args[1])
	case "stop":
		StopCommand(guildID, vInstance)
	case "pause":
		PauseCommand(vInstance)
	case "skip":
		SkipCommand(vInstance)
	case "disconnect":
		DisconnectCommand(guildID, vInstance)
	case "radio":
		RadioCommand(vInstance)
	}
}

func ChVoiceJoin(guildID string, vInstance *VoiceInstance) error {
	voiceChannelID := SearchVoiceChannel(guildID, message.Author.ID)
	vConnect, err := session.ChannelVoiceJoin(guildID, voiceChannelID, false, false)
	if err != nil {
		log.Println("Error: channel voice join", err)
		return err
	}
	vInstance.voice = vConnect
	return nil
}

func SearchGuild() (guildID string, err error) {
	channel, err := session.Channel(message.ChannelID)
	if err != nil {
		return
	}
	guildID = channel.GuildID
	return
}

func SearchVoiceChannel(guildID string, userID string) (voiceChannelID string) {
	for _, g := range session.State.Guilds {
		if g.ID != guildID {
			continue
		}

		for _, v := range g.VoiceStates {
			if v.UserID == userID {
				return v.ChannelID
			}
		}
	}
	return ""
}

// ChMessageSend send a message and auto-remove it in a time
func ChMessageSend(textChannelID, message string) {
	for i := 0; i < 10; i++ {
		msg, err := session.ChannelMessageSend(textChannelID, message)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		msgToPurgeQueue(msg)
		break
	}
}

// msgToPurgeQueue
func msgToPurgeQueue(m *discordgo.Message) {
	if o.DiscordPurgeTime > 0 {
		timestamp := time.Now().UTC().Unix()
		message := PurgeMessage{
			m.ID,
			m.ChannelID,
			timestamp,
		}
		purgeQueue = append(purgeQueue, message)
	}
}



// AddTimeDuration calculate the total time duration
func AddTimeDuration(t TimeDuration) (total TimeDuration) {
	total.Second = t.Second % 60
	t.Minute = t.Minute + t.Second/60
	total.Minute = t.Minute % 60
	t.Hour = t.Hour + t.Minute/60
	total.Hour = t.Hour % 24
	total.Day = t.Day + t.Hour/24
	return
}
