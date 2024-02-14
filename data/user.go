package data

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	RealName    string `json:"real_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
}
