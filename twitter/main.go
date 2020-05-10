package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/kelseyhightower/envconfig"
)

var hashTag = "#母の日"
var changed chan bool = make(chan bool)

type TwitterConfig struct {
	ConsumerKey       string `envconfig:"CONSUMER_KEY"`
	ConsumerSecret    string `envconfig:"CONSUMER_SECRET"`
	AccessToken       string `envconfig:"ACCESS_TOKEN"`
	AccessTokenSecret string `envconfig:"ACCESS_TOKEN_SECRET"`
}

func main() {
	go startTwitterStreaming()
	go changeHashTag()

	time.Sleep(20 * time.Second)
}

func changeHashTag() {
	hashTags := []string{"#メイドの日", "#青嵐実装決定", "#検察庁法改正案に抗議します"}

	for _, tag := range hashTags {
		time.Sleep(5 * time.Second)
		fmt.Printf("Changed to %s----------------------------------------------------\n", tag)
		hashTag = tag
		changed <- true
	}
}

func startTwitterStreaming() {
	if hashTag == "" {
		return
	}

	var c TwitterConfig
	envconfig.Process("TWITTER", &c)
	config := oauth1.NewConfig(c.ConsumerKey, c.ConsumerSecret)
	token := oauth1.NewToken(c.AccessToken, c.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		if time.Now().Unix()%3 == 0 {
			fmt.Println(tweet.Text)
		}
	}

	for {
		filterParams := &twitter.StreamFilterParams{Track: []string{hashTag}}
		stream, err := client.Streams.Filter(filterParams)
		if err != nil {
			log.Fatal(err)
		}

		go demux.HandleChan(stream.Messages)

		<-changed
		fmt.Println("Stop streaming*************************************************")
		stream.Stop()
	}
}
