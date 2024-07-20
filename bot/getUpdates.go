package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func GetUpdates(token string, offset int, timeout int) ([]Update, error) {
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
