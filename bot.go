package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
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
			fmt.Println("No updates for", time.Now())
			continue
		}
		for _, update := range updates {
			fmt.Println(update)
			offset = update.UpdateID + 1

			command := update.Message.Text

			var err error

			greetings, _ := regexp.MatchString(`hi|hey|hello|good (?:morning|day|evening|night)|yo|greetings|привет|приветствую|здравствуй|здравствуйте|даров|дарова|дроу|прив|хай|хеллоу|буп|доброе утро|добрый (?:день|вечер)|доброй ночи`, strings.ToLower(command))
			pic, _ := regexp.MatchString(`(?i)(дай пикчу)`, command)

			if greetings {
				_, err := sendMessage(update.Message.Chat.ID, "буп\n!")
				if err != nil {
					fmt.Println(err.Error())
				}
			} else if pic {
				picPath, err := randomPicPath("/home/zoob/Documents/gulty-pleasure/zoobot/pics")
				if err != nil {
					fmt.Println(err.Error())
				}
				_, err = sendPhoto(update.Message.Chat.ID, picPath)
				if err != nil {
					fmt.Println(err.Error())
				}
			} else {
				_, err := sendMessage(update.Message.Chat.ID, "я хз что ты от меня хочешь...")
				if err != nil {
					fmt.Println(err.Error())
				}
			}

			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

	//fmt.Println(u.String())
}

func randomPicPath(path string) (string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}
	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(files))
	randomFileName := files[randomIndex].Name()

	fullPath := filepath.Join(path, randomFileName)

	return fullPath, nil
}

func rawRequest(httpMethod string, url string, body io.Reader) ([]byte, error) {
	resp, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func getUpdates(token string, offset int, timeout int) ([]Update, error) {
	q := url.Values{}
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
		Path:     fmt.Sprint("bot", os.Getenv("BOT_TOKEN"), "/sendMessage"),
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

func sendPhoto(chatID int, filePath string) (message Message, err error) {
	var u = url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
		Path:   fmt.Sprint("bot", os.Getenv("BOT_TOKEN"), "/sendPhoto"),
	}

	file, err := os.Open(filePath)
	if err != nil {
		return Message{}, err
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	log.Println(filepath.Base(filePath))
	chatIDPart, err := writer.CreateFormField("chat_id")
	if err != nil {
		return Message{}, err
	}
	_, err = chatIDPart.Write([]byte(strconv.Itoa(chatID)))
	if err != nil {
		return Message{}, err
	}
	contentPart, err := writer.CreateFormFile("photo", filepath.Base(filePath))
	if err != nil {
		return Message{}, err
	}
	_, err = io.Copy(contentPart, file)
	if err != nil {
		return Message{}, err
	}
	err = writer.Close()
	if err != nil {
		return Message{}, err
	}

	req, err := http.NewRequest("POST", u.String(), body)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	res := &http.Client{}
	resp, err := res.Do(req)
	if err != nil {
		return Message{}, err
	}

	resBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return Message{}, fmt.Errorf("status code %d. body %s", resp.StatusCode, resBody)
	}

	return Message{}, nil
}
