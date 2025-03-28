package bot

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func SendVoice(chatID int, voicePath string) (Message, error) {
	var u = url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
		Path:   fmt.Sprint("bot", os.Getenv("BOT_TOKEN"), "/sendVoice"),
	}

	file, err := os.Open(voicePath)
	if err != nil {
		return Message{}, err
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	chatIDPart, err := writer.CreateFormField("chat_id")
	if err != nil {
		return Message{}, err
	}
	_, err = chatIDPart.Write([]byte(strconv.Itoa(chatID)))
	if err != nil {
		return Message{}, err
	}
	contentPart, err := writer.CreateFormFile("voice", filepath.Base(voicePath))
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
