package data

type ChatRead struct {
	ChatID    int    `json:"chat_id"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type ChatRoom struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type Chat struct {
	ID        int    `json:"id"`
	RoomID    int    `json:"room_id"`
	SenderID  int    `json:"sender_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
