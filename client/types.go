package client

type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}
type From struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}
type Message struct {
	ID   int    `json:"message_id"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
	Date int    `json:"date"`
	Text string `json:"text"`
}
type Update struct {
	ID      int     `json:"update_id"`
	Message Message `json:"message"`
}

type GetUpdatesReturn struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type SendMessageReturn struct {
	Ok bool `json:"ok"`
}
