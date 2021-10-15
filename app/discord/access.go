package discord

import (
	"strings"
	"os"
)

var (
	RoleList []string
)

func isAccess(memberRoleList []string) bool {
	if len(RoleList) == 0 {
		RoleList =  strings.Split(os.Getenv("ROLE_LIST"), ":")
	}

	for _, memberRole := range memberRoleList {
		for _, role := range RoleList {
			if memberRole == role {
				return true
			}
		}
	}
	return false
}
