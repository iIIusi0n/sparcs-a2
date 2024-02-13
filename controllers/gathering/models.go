package gathering

type Gathering struct {
	ID              int     `json:"id"`
	UserID          int     `json:"user_id"`
	Title           string  `json:"title"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	DateTime        string  `json:"time"`
	DurationInHours int     `json:"duration_in_hours"`
	MaxParticipants int     `json:"max_participants"`
	ThumbnailURL    string  `json:"thumbnail_url"`
	CreatedAt       string  `json:"created_at"`
}

type Participant struct {
	ID          int    `json:"id"`
	GatheringID int    `json:"gathering_id"`
	UserID      int    `json:"user_id"`
	CreatedAt   string `json:"created_at"`
}

type Location struct {
	ID          int    `json:"id"`
	GatheringID int    `json:"gathering_id"`
	PostID      int    `json:"post_id"`
	CreatedAt   string `json:"created_at"`
}

const (
	SortOptionLatest = "latest"
	SortOptionNearby = "nearby"
	SortOptionLiked  = "liked"
)
