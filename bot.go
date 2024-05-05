package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// https://api.telegram.org/bot<token>/METHOD_NAME

type Bot struct {
	Token string `json:"token"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func bot() Bot {
	bot := Bot{
		Token: os.Getenv("BOT_TOKEN"),
	}
	return bot
}

func (bot *Bot) SendMessage(target int, text string) {

}

func main() {

}
