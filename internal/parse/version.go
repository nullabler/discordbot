package parse

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/unixoff/discordbot/internal/discord"
)

const PHP_CHANNEL_ID = "873187094380576918"
const RUST_CHANNEL_ID = "835447565666746408"
const GO_CHANNEL_ID = "873133266629165076"

const RUST_URL = "https://github.com/rust-lang/rust/blob/master/RELEASES.md"
const GO_URL = "https://go.dev/doc/devel/release"

func Route(d *discord.Discord) {
	switch true {
	case d.HasChannelID(PHP_CHANNEL_ID):
		list := lastVersionPHP()
		for _, item := range list {
			d.MessageSend("Current Stable: v" + item[1] + " https://php.net" + item[2])
		}

	case d.HasChannelID(RUST_CHANNEL_ID):
		item := lastVersionRust()
		d.MessageSend("Current: v" + item[2] + " " + RUST_URL + item[1])

	case d.HasChannelID(GO_CHANNEL_ID):
		item := lastVersionGo()
		d.MessageSend("Current: v" + item[1] + " (" + item[2] + ") " + "https://go.dev/doc/go" + item[1])
	}
}

func lastVersionPHP() [][]string {
	resp, err := http.Get("https://www.php.net/downloads.php")
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	body := string(bodyByte)
	reg := regexp.MustCompile(`Current\s+Stable[^>]+>\s+PHP\s+([\d\.]+)[^<]+<.+href="([^"]+)`)
	return reg.FindAllStringSubmatch(body, -1)
}

func lastVersionRust() []string {
	resp, err := http.Get(RUST_URL)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	body := string(bodyByte)
	reg := regexp.MustCompile(`href="([^"]+)[^>]+>Version\s([\d\.]+\s\([\d\-]+\))`)
	return reg.FindStringSubmatch(body)
}

func lastVersionGo() []string {
	resp, err := http.Get(GO_URL)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	body := string(bodyByte)
	reg := regexp.MustCompile(`>go([\d\.]+)\s\(released\s([\d\-]+)\)`)
	return reg.FindStringSubmatch(body)
}
