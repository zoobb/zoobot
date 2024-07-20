// WIP
// todo rawRequest

package bot

import (
	"io"
	"net/http"
)

func RawRequest(httpMethod string, url string, body io.Reader) ([]byte, error) {
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
