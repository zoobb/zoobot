package main

import (
	"encoding/json"
	"fmt"
	"github.com/ZubovSL/zoobot/types"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	//fmt.Println("Hello World")

	// https://api.telegram.org/bot<token>/METHOD_NAME
	botURL := "https://api.telegram.org/bot" + os.Getenv("BOT_TOKEN")
	fmt.Println(getUpdates(botURL))
}

func getUpdates(botURL string) ([]types.Update, error) {
	resp, err := http.Get(botURL + "/getUpdates")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response types.Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return response.Result, nil
}
