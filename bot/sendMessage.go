package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func SendMessage(chatID int, text string) (Message, error) {
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
