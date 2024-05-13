package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// Update https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID          int     `json:"update_id"`
	Message           Message `json:"message"`
	EditedMessage     Message `json:"edited_message"`
	ChannelPost       Message `json:"channel_post"`
	EditedChannelPost Message `json:"edited_channel_post"`
}

// Message https://core.telegram.org/bots/api#message
type Message struct {
	MessageID       int    `json:"message_id"`
	MessageThreadID int    `json:"message_thread_id"`
	From            Chat   `json:"from"`
	SenderChat      Chat   `json:"sender_chat"`
	Date            int    `json:"date"`
	Chat            Chat   `json:"chat"`
	Text            string `json:"text"`
}

// Chat https://core.telegram.org/bots/api#chat
type Chat struct {
	ID              int      `json:"id"`
	Type            string   `json:"type"`
	Title           string   `json:"title"`
	Username        string   `json:"username"`
	FirstName       string   `json:"first_name"`
	LastName        string   `json:"last_name"`
	IsForum         bool     `json:"is_forum"`
	ActiveUsernames []string `json:"active_usernames"`
}

type Response struct {
	Result []Update `json:"result"`
}

// https://api.telegram.org/bot<token>/METHOD_NAME

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	var offset = 0
	for {
		updates, err := getUpdates(os.Getenv("BOT_TOKEN"), offset, 60)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if len(updates) == 0 {
			fmt.Println("No updates...")
			continue
		}
		for _, update := range updates {
			offset = update.UpdateID + 1

			// Echo
			/* _, err := sendMessage(update.Message.Chat.ID, update.Message.Text)
			      if err != nil {
			   			return
			      }
			      fmt.Println(update.Message.Text) */

			command := update.Message.Text

			switch command {
			case "дарова":
				_, err := sendMessage(update.Message.Chat.ID, "буп")
				if err != nil {
					fmt.Println(err.Error())
				}
			default:
				_, err := sendMessage(update.Message.Chat.ID, "я хз что ты от меня хочешь...")
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}

	//fmt.Println(u.String())
}

func getUpdates(token string, offset int, timeout int) ([]Update, error) {
	var q = make(url.Values)
	q.Set("offset", strconv.Itoa(offset))
	q.Set("timeout", strconv.Itoa(timeout))

	var u = url.URL{
		Scheme:   "https",
		Host:     "api.telegram.org",
		Path:     fmt.Sprintf("bot%s/%s", token, "getUpdates"),
		RawQuery: q.Encode(),
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	//resp, err := http.DefaultClient.Do(req)
	//fmt.Println(resp, err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r Response
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r.Result, nil
}

func sendMessage(chatID int, text string) (Message, error) {
	var q = make(url.Values)
	q.Set("chat_id", strconv.Itoa(chatID))
	q.Set("text", text)

	var u = url.URL{
		Scheme:   "https",
		Host:     "api.telegram.org",
		Path:     fmt.Sprintf("bot%s/%s", os.Getenv("BOT_TOKEN"), "sendMessage"),
		RawQuery: q.Encode(),
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return Message{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, err
	}

	var r Message
	if err := json.Unmarshal(body, &r); err != nil {
		return Message{}, err
	}

	return r, nil
}
