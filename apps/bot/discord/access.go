package discord

import (
	"os"
	"strings"
)

const SEPARATION = ":"

var (
	RoleList []string
)

func init() {
	if len(RoleList) == 0 {
		RoleList = strings.Split(os.Getenv("ROLE_LIST"), SEPARATION)
	}
}

func isAccess() bool {
	for _, memberRole := range message.Member.Roles {
		for _, role := range RoleList {
			if memberRole == role {
				return true
			}
		}
	}

	return false
}

func isAccessForMusic() bool {
	if os.Getenv("MUSIC_CHANNEL") != message.ChannelID{
		return false
	}

	return true
}
