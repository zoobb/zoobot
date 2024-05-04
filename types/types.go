package types

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
