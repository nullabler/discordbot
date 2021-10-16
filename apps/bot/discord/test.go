package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func testHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.MessageReactionAdd(message.ChannelID, message.ID, "✅")


	// fmt.Println(message.Author.ID)
	fmt.Printf("%#v \n", SearchVoiceChannel(session, message.Author.ID))
	// voiceState, _ := session.State.VoiceState(message.GuildID, message.Author.ID)
	// fmt.Printf("%#v \n", voiceState.ChannelID)
	// session.ChannelVoiceJoin(message.GuildID, message.ChannelID, false, false)


	if session.VoiceReady {
		session.ChannelMessageSend(message.ChannelID, "test")
	}
}

func SearchVoiceChannel(session *discordgo.Session, user string) (voiceChannelID string) {

	fmt.Println(session.State.Guilds[0].ApplicationID)
	for _, g := range session.State.Guilds {
		for _, v := range g.VoiceStates {
			if v.UserID == user {
				return v.ChannelID
			}
		}
	}
	return ""
}
