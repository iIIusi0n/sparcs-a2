package post

type Post struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	ImageURL    string  `json:"image_url"`
	HashTags    string  `json:"hash_tags"`
	CreatedAt   string  `json:"created_at"`
}

type Like struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	PostID    int    `json:"post_id"`
	CreatedAt string `json:"created_at"`
}
