package main

import (
	"encoding/json"
	"fmt"
	"github.com/ZubovSL/zoobot/types"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func Polling(delayMS int) types.Update {
	offset := 0
	for {
		updates, err := getUpdates(offset)
		if err != nil {
			log.Fatal(err)
		}

		for _, update := range updates {
			offset = update.UpdateID + 1
			return update
		}
		time.Sleep(time.Duration(delayMS))
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
