package data

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profile_picture"`
	CreatedAt      string `json:"created_at"`
}

func (m *manager) CreateUser(user *User) (int, error) {
	result, err := m.db.Exec("INSERT INTO users (username, name, profile_picture) VALUES (?, ?, ?)", user.Username, user.Name, user.ProfilePicture)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *manager) ReadUser(id int) (*User, error) {
	row := m.db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	var u User
	err := row.Scan(&u.ID, &u.Username, &u.Name, &u.ProfilePicture, &u.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (m *manager) UpdateUser(user *User) error {
	_, err := m.db.Exec("UPDATE users SET username = ?, name = ?, profile_picture = ? WHERE id = ?", user.Username, user.Name, user.ProfilePicture, user.ID)
	return err
}

func (m *manager) DeleteUser(id int) error {
	_, err := m.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
