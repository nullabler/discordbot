package music

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var o = &Options{}

func SearchGuild(textChannelID string) (guildID string) {
	channel, _ := session.Channel(textChannelID)
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
