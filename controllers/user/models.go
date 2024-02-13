package user

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profile_picture"`
	CreatedAt      string `json:"created_at"`
}
