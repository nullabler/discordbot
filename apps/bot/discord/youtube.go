package discord

import (
	"log"

	"github.com/kkdai/youtube/v2"
)

type Youtube struct {
	state  *State
	client youtube.Client
}

func newYoutube() *Youtube {
	return &Youtube{
		client: youtube.Client{},
	}
}

func (youtube *Youtube) init(state *State) {
	youtube.state = state
}

func (youtube *Youtube) find(search string) (pkgSong PkgSong, err error) {
	song := Song{
		ChannelID: youtube.state.message.ChannelID,
		User:      youtube.state.message.Author.ID,
	}

	// playlist, _ := client.GetPlaylist(search)
	// for _, item := range playlist.Videos {
	// 	log.Println(item)
	// }

	video, err := youtube.client.GetVideo(search)
	if err != nil {
		log.Println("Fatal: Youtube get video", err)
		return
	}

	formats := video.Formats.WithAudioChannels()
	streamURL, err := youtube.client.GetStreamURL(video, &formats[0])
	if err != nil {
		log.Println("Fatal: Youtube get stream URL", err)
		return
	}
	// log.Println(streamURL)

	song.ID = video.Author
	song.VideoURL = streamURL
	song.VidID = video.ID
	song.Title = video.Title
	song.Duration = video.Duration.String()

	pkgSong = PkgSong{
		data: song,
	}

	return
}
