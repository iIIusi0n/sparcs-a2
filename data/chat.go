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

func (m *manager) CreateChatRoom(chatRoom *ChatRoom) (int, error) {
	query := "INSERT INTO chat_rooms (name) VALUES (?)"
	result, err := m.db.Exec(query, chatRoom.Name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *manager) ReadChatRoom(id int) (*ChatRoom, error) {
	query := "SELECT * FROM chat_rooms WHERE id = ?"
	row := m.db.QueryRow(query, id)

	chatRoom := ChatRoom{}
	err := row.Scan(&chatRoom.ID, &chatRoom.Name, &chatRoom.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &chatRoom, nil
}

func (m *manager) ReadChatRooms() ([]*ChatRoom, error) {
	query := "SELECT * FROM chat_rooms"
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}

	chatRooms := []*ChatRoom{}
	for rows.Next() {
		chatRoom := ChatRoom{}
		err := rows.Scan(&chatRoom.ID, &chatRoom.Name, &chatRoom.CreatedAt)
		if err != nil {
			return nil, err
		}

		chatRooms = append(chatRooms, &chatRoom)
	}

	return chatRooms, nil
}

func (m *manager) UpdateChatRoom(chatRoom *ChatRoom) error {
	query := "UPDATE chat_rooms SET name = ? WHERE id = ?"
	_, err := m.db.Exec(query, chatRoom.Name, chatRoom.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *manager) DeleteChatRoom(id int) error {
	query := "DELETE FROM chat_rooms WHERE id = ?"
	_, err := m.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *manager) CreateChat(chat *Chat) (int, error) {
	query := "INSERT INTO chats (room_id, sender_id, content) VALUES (?, ?, ?)"
	result, err := m.db.Exec(query, chat.RoomID, chat.SenderID, chat.Content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *manager) ReadChat(id int) (*Chat, error) {
	query := "SELECT * FROM chats WHERE id = ?"
	row := m.db.QueryRow(query, id)

	chat := Chat{}
	err := row.Scan(&chat.ID, &chat.RoomID, &chat.SenderID, &chat.Content, &chat.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &chat, nil
}

func (m *manager) ReadChatsByRoomID(roomID int) ([]*Chat, error) {
	query := "SELECT * FROM chats WHERE room_id = ? ORDER BY created_at"
	rows, err := m.db.Query(query, roomID)
	if err != nil {
		return nil, err
	}

	chats := []*Chat{}
	for rows.Next() {
		chat := Chat{}
		err := rows.Scan(&chat.ID, &chat.RoomID, &chat.SenderID, &chat.Content, &chat.CreatedAt)
		if err != nil {
			return nil, err
		}

		chats = append(chats, &chat)
	}

	return chats, nil
}

func (m *manager) UpdateChat(chat *Chat) error {
	query := "UPDATE chats SET room_id = ?, sender_id = ?, content = ? WHERE id = ?"
	_, err := m.db.Exec(query, chat.RoomID, chat.SenderID, chat.Content, chat.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *manager) DeleteChat(id int) error {
	query := "DELETE FROM chats WHERE id = ?"
	_, err := m.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *manager) CreateChatRead(chatRead *ChatRead) error {
	query := "INSERT INTO chat_read (chat_id, user_id) VALUES (?, ?)"
	_, err := m.db.Exec(query, chatRead.ChatID, chatRead.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (m *manager) ReadChatRead(chatID, userID int) (*ChatRead, error) {
	query := "SELECT * FROM chat_read WHERE chat_id = ? AND user_id = ?"
	row := m.db.QueryRow(query, chatID, userID)

	chatRead := ChatRead{}
	err := row.Scan(&chatRead.ChatID, &chatRead.UserID, &chatRead.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &chatRead, nil
}

func (m *manager) UpdateChatRead(chatRead *ChatRead) error {
	query := "UPDATE chat_read SET chat_id = ?, user_id = ? WHERE chat_id = ? AND user_id = ?"
	_, err := m.db.Exec(query, chatRead.ChatID, chatRead.UserID, chatRead.ChatID, chatRead.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (m *manager) DeleteChatRead(chatID, userID int) error {
	query := "DELETE FROM chat_read WHERE chat_id = ? AND user_id = ?"
	_, err := m.db.Exec(query, chatID, userID)
	if err != nil {
		return err
	}

	return nil
}
