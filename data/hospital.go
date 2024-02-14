package data

type Hospital struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	NumberOfDoctor int     `json:"number_of_doctor"`
	Address        string  `json:"address"`
	PhoneNumber    string  `json:"phone_number"`
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

func (m *manager) CreateHospital(hospital *Hospital) (int, error) {
	query := "INSERT INTO hospitals (name, latitude, longitude, number_of_doctor) VALUES (?, ?, ?, ?)"
	result, err := m.db.Exec(query, hospital.Name, hospital.Latitude, hospital.Longitude, hospital.NumberOfDoctor)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *manager) ReadHospital(id int) (*Hospital, error) {
	query := "SELECT * FROM hospitals WHERE id = ?"
	row := m.db.QueryRow(query, id)

	hospital := Hospital{}
	err := row.Scan(&hospital.ID, &hospital.Name, &hospital.Latitude, &hospital.Longitude, &hospital.NumberOfDoctor, &hospital.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &hospital, nil
}

func (m *manager) ReadHospitals() ([]*Hospital, error) {
	query := "SELECT * FROM hospitals"
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	hospitals := []*Hospital{}
	for rows.Next() {
		hospital := Hospital{}
		err := rows.Scan(&hospital.ID, &hospital.Name, &hospital.Latitude, &hospital.Longitude, &hospital.NumberOfDoctor, &hospital.CreatedAt)
		if err != nil {
			return nil, err
		}
		hospitals = append(hospitals, &hospital)
	}

	return hospitals, nil

}

func (m *manager) UpdateHospital(hospital *Hospital) error {
	query := "UPDATE hospitals SET name = ?, latitude = ?, longitude = ?, number_of_doctor = ? WHERE id = ?"
	_, err := m.db.Exec(query, hospital.Name, hospital.Latitude, hospital.Longitude, hospital.NumberOfDoctor, hospital.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *manager) DeleteHospital(id int) error {
	query := "DELETE FROM hospitals WHERE id = ?"
	_, err := m.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *manager) CreateInHospital(inHospital *InHospital) error {
	query := "INSERT INTO in_hospitals (hospital_id, user_id) VALUES (?, ?)"
	_, err := m.db.Exec(query, inHospital.HospitalID, inHospital.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (m *manager) ReadInHospital(hospitalID, userID int) (*InHospital, error) {
	query := "SELECT * FROM in_hospitals WHERE hospital_id = ? AND user_id = ?"
	row := m.db.QueryRow(query, hospitalID, userID)

	inHospital := InHospital{}
	err := row.Scan(&inHospital.HospitalID, &inHospital.UserID, &inHospital.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &inHospital, nil
}

func (m *manager) ReadInHospitalsByHospitalID(hospitalID int) ([]*InHospital, error) {
	query := "SELECT * FROM in_hospitals WHERE hospital_id = ? and created_at > DATE_SUB(NOW(), INTERVAL 1 DAY)"
	rows, err := m.db.Query(query, hospitalID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	inHospitals := []*InHospital{}
	for rows.Next() {
		inHospital := InHospital{}
		err := rows.Scan(&inHospital.HospitalID, &inHospital.UserID, &inHospital.CreatedAt)
		if err != nil {
			return nil, err
		}
		inHospitals = append(inHospitals, &inHospital)
	}

	return inHospitals, nil
}

func (m *manager) UpdateInHospital(inHospital *InHospital) error {
	query := "UPDATE in_hospitals SET hospital_id = ?, user_id = ? WHERE hospital_id = ? AND user_id = ?"
	_, err := m.db.Exec(query, inHospital.HospitalID, inHospital.UserID, inHospital.HospitalID, inHospital.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (m *manager) DeleteInHospital(hospitalID, userID int) error {
	query := "DELETE FROM in_hospitals WHERE hospital_id = ? AND user_id = ?"
	_, err := m.db.Exec(query, hospitalID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (m *manager) CreateWaitingNumber(waitingNumber *WaitingNumber) error {
	query := "INSERT INTO waiting_numbers (hospital_id, user_id, number) VALUES (?, ?, ?)"
	_, err := m.db.Exec(query, waitingNumber.HospitalID, waitingNumber.UserID, waitingNumber.Number)
	if err != nil {
		return err
	}

	return nil
}

func (m *manager) ReadWaitingNumber(hospitalID, userID int) (*WaitingNumber, error) {
	query := "SELECT * FROM waiting_numbers WHERE hospital_id = ? AND user_id = ?"
	row := m.db.QueryRow(query, hospitalID, userID)

	waitingNumber := WaitingNumber{}
	err := row.Scan(&waitingNumber.UserID, &waitingNumber.HospitalID, &waitingNumber.Number, &waitingNumber.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &waitingNumber, nil
}

func (m *manager) ReadLatestWaitingNumber(hospitalID int) (*WaitingNumber, error) {
	query := "SELECT * FROM waiting_numbers WHERE hospital_id = ? ORDER BY created_at DESC LIMIT 1"
	row := m.db.QueryRow(query, hospitalID)

	waitingNumber := WaitingNumber{}
	err := row.Scan(&waitingNumber.UserID, &waitingNumber.HospitalID, &waitingNumber.Number, &waitingNumber.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &waitingNumber, nil
}

func (m *manager) UpdateWaitingNumber(waitingNumber *WaitingNumber) error {
	query := "UPDATE waiting_numbers SET hospital_id = ?, user_id = ?, number = ? WHERE hospital_id = ? AND user_id = ?"
	_, err := m.db.Exec(query, waitingNumber.HospitalID, waitingNumber.UserID, waitingNumber.Number, waitingNumber.HospitalID, waitingNumber.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (m *manager) DeleteWaitingNumber(hospitalID, userID int) error {
	query := "DELETE FROM waiting_numbers WHERE hospital_id = ? AND user_id = ?"
	_, err := m.db.Exec(query, hospitalID, userID)
	if err != nil {
		return err
	}

	return nil
}
