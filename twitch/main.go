package main

import (
	"crypto/tls"
	"fmt"
	"time"

	irc "github.com/thoj/go-ircevent"

	"github.com/kelseyhightower/envconfig"
)

const serverssl = "irc.chat.twitch.tv:6697"

var twitchChannel = "#mogra"
var changed chan bool = make(chan bool)

type TwitchConfig struct {
	Nick     string
	Password string
}

func main() {
	go startTwitterStreaming()
	go changeTwitchChannel()

	time.Sleep(20 * time.Second)
}

func changeTwitchChannel() {
	hashTags := []string{"#hinas3", "#twitch_yuce", "#sooflower"}

	for _, tag := range hashTags {
		time.Sleep(10 * time.Second)
		fmt.Printf("Changed to %s----------------------------------------------------\n", tag)
		twitchChannel = tag
		changed <- true
	}
}

func startTwitterStreaming() {
	if twitchChannel == "" {
		return
	}

	var config TwitchConfig
	envconfig.Process("TWITCH", &config)

	nick := config.Nick
	con := irc.IRC(nick, nick)

	con.Password = config.Password
	con.UseTLS = true
	con.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	con.AddCallback("001", func(e *irc.Event) { con.Join(twitchChannel) })
	con.AddCallback("PRIVMSG", func(e *irc.Event) {
		fmt.Println(e.Message() + ": " + twitchChannel)
	})
	err := con.Connect(serverssl)
	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}

	go con.Loop()

	for {

		<-changed
		fmt.Println("Stop connection.................")
		con.Part()
	}
}
