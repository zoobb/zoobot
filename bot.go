package main

import (
	"fmt"
	"github.com/ZubovSL/zoobot/requests"
	"github.com/joho/godotenv"
	"log"
)

// https://api.telegram.org/bot<token>/METHOD_NAME

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	update := Polling(10)
	resp, err := requests.SendMessage(update.Message.Chat.ID, "boop")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
