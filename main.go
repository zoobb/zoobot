package main

import (
	"github.com/ZubovSL/zoobot/bot"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func main() {
	botToken := os.Getenv("BOT_TOKEN")
	picsFolder := os.Getenv("PICS_FOLDER")

	err := bot.NewBot(botToken, 60, picsFolder)
	if err != nil {
		return
	}

}
