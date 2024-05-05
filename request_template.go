package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// https://api.telegram.org/bot<token>/METHOD_NAME

func request_template(method string, data map[string]string) error {

	var u = url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
		Path:   fmt.Sprintf("bot%s/%s", os.Getenv("BOT_TOKEN"), method),
	}
	var q = u.Query()
	for key, value := range data {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}
