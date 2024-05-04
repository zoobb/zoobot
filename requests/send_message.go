package requests

import (
  "fmt"
  "github.com/ZubovSL/zoobot/types"
  "io"
  "net/http"
  "os"
)

func SendMessage(target int, message string) (msg types.Message, error) {
  resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s", os.Getenv("BOT_TOKEN"), target, message))
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  body, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

}
