package discord

const (
	c_DEFAULT             = 0
	c_AQUA                = 1752220
	c_GREEN               = 3066993
	c_BLUE                = 3447003
	c_PURPLE              = 10181046
	c_GOLD                = 15844367
	c_ORANGE              = 15105570
	c_RED                 = 15158332
	c_GREY                = 9807270
	c_DARKER_GREY         = 8359053
	c_NAVY                = 3426654
	c_DARK_AQUA           = 1146986
	c_DARK_GREEN          = 2067276
	c_DARK_BLUE           = 2123412
	c_DARK_PURPLE         = 7419530
	c_DARK_GOLD           = 12745742
	c_DARK_ORANGE         = 11027200
	c_DARK_RED            = 10038562
	c_DARK_GREY           = 9936031
	c_LIGHT_GREY          = 12370112
	c_DARK_NAVY           = 2899536
	c_LUMINOUS_VIVID_PINK = 16580705
	c_DARK_VIVID_PINK     = 12320855
)

func errorMessage(title string, message string) string {
	return "❌  **" + title + "**\n" + message
}

func successMessage(title string, message string) string {
	return "✅  **" + title + "**\n" + message
}

func permissionDeniedMessage() {
	session.ChannelMessageSend(message.ChannelID, errorMessage("Error", "Permission denied"))
}
