package bot

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func SendPhoto(chatID int, filePath string) (message Message, err error) {
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
