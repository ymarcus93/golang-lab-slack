package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/context"

	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/nlopes/slack"
	"github.com/sparrc/go-ping"
)

func main() {
	Listen()
}

func CheckServer(ip string) bool {
	fmt.Println("DEBUG: Passed IP: ", ip)

	pinger, err := ping.NewPinger(ip)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	if stats != nil {
		return true
	} else {
		return false
	}
}

func Listen() {
	bot := slackbot.New(os.Getenv("SLACK_API_TOKEN"))
	bot.Hear("server (.*)").MessageHandler(RespondRequestHandler)
	bot.Run()
}

func RespondRequestHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	// TODO: Respond to a `server <IP>` request
	fmt.Println("DEBUG: User sent: ", evt.Msg.Text)
	arrOfWords := strings.Fields(evt.Msg.Text)
	ip := arrOfWords[len(arrOfWords)-1]
	err := CheckServer(ip)
	if err == true {
		bot.Reply(evt, fmt.Sprintf("ONLINE"), slackbot.WithTyping)
	} else {
		bot.Reply(evt, fmt.Sprintf("OFFLINE"), slackbot.WithTyping)
	}
}
