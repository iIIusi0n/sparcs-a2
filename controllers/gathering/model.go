package gathering

type Gathering struct {
	ID              int     `json:"id"`
	UserID          int     `json:"user_id"`
	Title           string  `json:"title"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Time            string  `json:"time"`
	MaxParticipants int     `json:"max_participants"`
	CreatedAt       string  `json:"created_at"`
}

type Participant struct {
	ID          int    `json:"id"`
	GatheringID int    `json:"gathering_id"`
	UserID      int    `json:"user_id"`
	CreatedAt   string `json:"created_at"`
}
