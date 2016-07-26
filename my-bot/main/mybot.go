package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/context"

	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/briandowns/openweathermap"
	"github.com/nlopes/slack"
)

func main() {
	ListenForWeather()
}

func GetWeather(place string) (*openweathermap.CurrentWeatherData, error) {
	// TODO: get the weather for the place
	return nil, nil
}

func ListenForWeather() {
	bot := slackbot.New(os.Getenv("SLACK_API_TOKEN"))
	// TODO: add a bot.Hear listener here for a `weather <place>` request
	bot.Run()
}

func WeatherHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	// TODO: Respond to a `weather <place>` request
}
