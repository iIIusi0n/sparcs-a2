package data

import (
	"api-server/controllers/post"
)

func (m *manager) CreatePost(post *post.Post) (int, error) {
	result, err := m.db.Exec("INSERT INTO posts (user_id, title, description, latitude, longitude, image_url, hashtags) VALUES (?, ?, ?, ?, ?, ?, ?)", post.UserID, post.Title, post.Description, post.Latitude, post.Longitude, post.ImageURL, post.HashTags)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *manager) ReadPost(id int) (*post.Post, error) {
	row := m.db.QueryRow("SELECT * FROM posts WHERE id = ?", id)

	var p post.Post
	err := row.Scan(&p.ID, &p.UserID, &p.Title, &p.Description, &p.Latitude, &p.Longitude, &p.ImageURL, &p.HashTags, &p.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (m *manager) UpdatePost(post *post.Post) error {
	_, err := m.db.Exec("UPDATE posts SET title = ?, description = ?, latitude = ?, longitude = ?, image_url = ?, hashtags = ? WHERE id = ?", post.Title, post.Description, post.Latitude, post.Longitude, post.ImageURL, post.HashTags, post.ID)
	return err
}

func (m *manager) DeletePost(id int) error {
	_, err := m.db.Exec("DELETE FROM posts WHERE id = ?", id)
	return err
}

func (m *manager) CreateLike(like *post.Like) (int, error) {
	result, err := m.db.Exec("INSERT INTO likes (user_id, post_id) VALUES (?, ?)", like.UserID, like.PostID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *manager) ReadLike(id int) (*post.Like, error) {
	row := m.db.QueryRow("SELECT * FROM likes WHERE id = ?", id)

	var l post.Like
	err := row.Scan(&l.ID, &l.UserID, &l.PostID, &l.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &l, nil
}

func (m *manager) UpdateLike(like *post.Like) error {
	_, err := m.db.Exec("UPDATE likes SET user_id = ?, post_id = ? WHERE id = ?", like.UserID, like.PostID, like.ID)
	return err
}

func (m *manager) DeleteLike(id int) error {
	_, err := m.db.Exec("DELETE FROM likes WHERE id = ?", id)
	return err
}

func (m *manager) CountLikeOnPost(postID int) (int, error) {
	row := m.db.QueryRow("SELECT COUNT(*) FROM likes WHERE post_id = ?", postID)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (m *manager) CheckUserLikedPost(userID, postID int) (bool, error) {
	row := m.db.QueryRow("SELECT COUNT(*) FROM likes WHERE user_id = ? AND post_id = ?", userID, postID)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (m *manager) ReadPostsByUserID(userID int) ([]*post.Post, error) {
	rows, err := m.db.Query("SELECT * FROM posts WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*post.Post
	for rows.Next() {
		var p post.Post
		err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Description, &p.Latitude, &p.Longitude, &p.ImageURL, &p.HashTags, &p.CreatedAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, &p)
	}

	return posts, nil
}

func (m *manager) CountPostsOnUser(userID int) (int, error) {
	row := m.db.QueryRow("SELECT COUNT(*) FROM posts WHERE user_id = ?", userID)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
