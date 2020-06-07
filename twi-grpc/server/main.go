package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/kelseyhightower/envconfig"

	pb "github.com/alitaso345/zatsu/twi-grpc/proto"
	"google.golang.org/grpc"
)

type TwitterConfig struct {
	ConsumerKey       string `envconfig:"CONSUMER_KEY"`
	ConsumerSecret    string `envconfig:"CONSUMER_SECRET"`
	AccessToken       string `envconfig:"ACCESS_TOKEN"`
	AccessTokenSecret string `envconfig:"ACCESS_TOKEN_SECRET"`
}

type timelineService struct{}

func (s *timelineService) Connect(req *pb.Room, stream pb.Timeline_ConnectServer) error {
	done := make(chan interface{})
	tc := generateTwitchCh(done, req)
	for {
		select {
		case msg := <-tc:
			err := stream.Send(&pb.Comment{Name: msg.User.ScreenName, Message: msg.Text, PlatformType: pb.PlatformType_TWITTER})
			if err != nil {
				log.Println("sending error")
				close(done)
				return err
			}
		}
	}
}

func generateTwitchCh(done <-chan interface{}, req *pb.Room) <-chan *twitter.Tweet {
	ch := make(chan *twitter.Tweet)
	go func() {
		defer func() {
			log.Println("close twitter ch")
			close(ch)
		}()
		var c TwitterConfig
		envconfig.Process("TWITTER", &c)
		config := oauth1.NewConfig(c.ConsumerKey, c.ConsumerSecret)
		token := oauth1.NewToken(c.AccessToken, c.AccessTokenSecret)
		httpClient := config.Client(oauth1.NoContext, token)

		client := twitter.NewClient(httpClient)

		demux := twitter.NewSwitchDemux()
		demux.Tweet = func(tweet *twitter.Tweet) {
			if tweet.RetweetedStatus != nil {
				return
			}
			fmt.Println(fmt.Sprintf("%s\n", tweet.Text))
			ch <- tweet
		}

		filterParams := &twitter.StreamFilterParams{Track: []string{req.GetHashTag()}}
		twitterStream, err := client.Streams.Filter(filterParams)
		if err != nil {
			log.Fatal(err)
		}
		defer twitterStream.Stop()

		go demux.HandleChan(twitterStream.Messages)
		<-done
		return
	}()

	return ch
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen port %v", port)
	}
	server := grpc.NewServer()
	pb.RegisterTimelineServer(server, &timelineService{})
	log.Printf("start server on port %s\n", port)
	server.Serve(lis)
}
