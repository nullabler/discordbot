package config

import (
	"os"
	"strings"
)

type Config struct {
	DiscordToken string
	YoutubeToken string
	CommandTarget string
	AccessRoleList []string
	MusicChannelList []string
}

func New() *Config {
	return &Config{
		DiscordToken:     os.Getenv("DISCORD_TOKEN"),
		YoutubeToken:     os.Getenv("YOUTUBE_TOKEN"),
		CommandTarget:    os.Getenv("DISCORD_TARGET"),
		AccessRoleList:   strings.Split(os.Getenv("ACCESS_ROLE_LIST"), SEPARATION),
		MusicChannelList: strings.Split(os.Getenv("MUSIC_CHANNEL_LIST"), SEPARATION),
	}
}
