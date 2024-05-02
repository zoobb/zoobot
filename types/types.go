package types

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageID int `json:"message_id"`
}

type Response struct {
	Result []Update `json:"result"`
}
