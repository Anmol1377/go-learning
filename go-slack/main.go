package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mmcdole/gofeed"
	"github.com/shomali11/slacker"
)

func PrintCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {

		fmt.Println("command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

type NewsHandler struct{}

func (h *NewsHandler) FetchNews() (string, error) {
	feedURL := "https://inc42.com/feed/"

	// Initialize gofeed parser
	fp := gofeed.NewParser()

	// Fetch and parse the RSS feed
	feed, err := fp.ParseURL(feedURL)
	if err != nil {
		return "", fmt.Errorf("could not fetch or parse feed: %v", err)
	}

	// Prepare a message with the latest news
	var newsMessage string
	if len(feed.Items) == 0 {
		newsMessage = "No recent news available at the moment."
	} else {
		newsMessage = "Here are the latest headlines:\n"
		for i, item := range feed.Items {
			if i >= 5 { 
				break
			}
			newsMessage += fmt.Sprintf("%d. *%s*\n%s\n\n", i+1, item.Title, item.Link)
		}
	}

	return newsMessage, nil
}

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-8386467779235-8372037795543-mLZpZmfzfq5Lt8JvDfYDCZAl")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A08BF2GBNDS-8386554356706-a1280d108a15843bbe0ca1e7745930d2e78a5faff05c59d0def0af79ed92016b")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go PrintCommandEvents(bot.CommandEvents())
	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(botCTX slacker.BotContext, Request slacker.Request, Response slacker.ResponseWriter) {
			Response.Reply("pong")
		},
	})
	bot.Command("hello", &slacker.CommandDefinition{
		Handler: func(botCTX slacker.BotContext, Request slacker.Request, Response slacker.ResponseWriter) {
			Response.Reply("hello, how are you?")
		},
	})


	newsHandler := &NewsHandler{}
	bot.Command("news", &slacker.CommandDefinition{
		Handler: func(botCTX slacker.BotContext, Request slacker.Request, Response slacker.ResponseWriter) {
			// Fetch the latest news
			news, err := newsHandler.FetchNews()
			if err != nil {
				// If there's an error, reply with an error message
				Response.Reply(fmt.Sprintf("Oops! Something went wrong: %v", err))
				return
			}

			// Send the latest news to the user
			Response.Reply(news)
		},
	})

	context, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(context)

	if err != nil {
		log.Fatal(err, "error here")
	}
}
