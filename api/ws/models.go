package ws

type Message struct {
	Content string `json:"content"`
}

type Response struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}
