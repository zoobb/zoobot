package bot

import (
	"log"
	"regexp"
	"strings"
	"time"
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

func NewBot(botToken string, pollingTimeoutSeconds int, picsFolder string) error {
	log.Println("ZOOBOT STARTED")
	var offset = 0
	for {
		updates, err := GetUpdates(botToken, offset, pollingTimeoutSeconds)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if len(updates) == 0 {
			log.Println("No updates for", time.Now())
			continue
		}

		// Updates loop
		for _, update := range updates {
			log.Println(update)
			offset = update.UpdateID + 1

			command := update.Message.Text

			var err error

			greetings, _ := regexp.MatchString(`hi|hey|hello|good (?:morning|day|evening|night)|yo|greetings`, strings.ToLower(command))
			pic, _ := regexp.MatchString(`(?i)(gimme a pic)`, command)

			if greetings {
				_, err := SendMessage(update.Message.Chat.ID, "hi\n!")
				if err != nil {
					log.Println(err.Error())
				}
			} else if pic {
				p, err := RandomPic(picsFolder)
				if err != nil {
					log.Println(err.Error())
				}
				_, err = SendPhoto(update.Message.Chat.ID, p)
				if err != nil {
					log.Println(err.Error())
				}
			} else {
				_, err := SendMessage(update.Message.Chat.ID, "i don't know what do you want from me...")
				if err != nil {
					log.Println(err.Error())
				}
			}

			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}
