package domain

type Notification struct {
	ID        uint   `json:"id"`
	Message   string `json:"message"`
	Info      string `json:"info"`
	CreatedAt string `json:"created_at"`
}
