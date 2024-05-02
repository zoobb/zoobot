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
	"time"
)

// https://api.telegram.org/bot<token>/METHOD_NAME

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	offset := 0
	for {
		updates, err := getUpdates(offset)
		if err != nil {
			log.Println(err)
		}

		for _, update := range updates {
			fmt.Println(update.Message.Text)
			offset = update.UpdateID + 1
		}
		time.Sleep(10)
	}
}

func getUpdates(offset int) ([]types.Update, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d", os.Getenv("BOT_TOKEN"), offset))
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
