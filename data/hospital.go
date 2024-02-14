package data

type Hospital struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	NumberOfDoctor int     `json:"number_of_doctor"`
	CreatedAt      string  `json:"created_at"`
}

type InHospital struct {
	HospitalID int    `json:"hospital_id"`
	UserID     int    `json:"user_id"`
	CreatedAt  string `json:"created_at"`
}

type WaitingNumber struct {
	UserID     int    `json:"user_id"`
	HospitalID int    `json:"hospital_id"`
	Number     int    `json:"number"`
	CreatedAt  string `json:"created_at"`
}
