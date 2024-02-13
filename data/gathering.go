package data

import (
	"api-server/controllers/gathering"
)

func (m *manager) CreateGathering(gathering *gathering.Gathering) (int, error) {
	result, err := m.db.Exec("INSERT INTO gatherings (user_id, title, latitude, longitude, date_time, max_participants) VALUES (?, ?, ?, ?, ?, ?)", gathering.UserID, gathering.Title, gathering.Latitude, gathering.Longitude, gathering.DateTime, gathering.MaxParticipants)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *manager) ReadGathering(id int) (*gathering.Gathering, error) {
	row := m.db.QueryRow("SELECT * FROM gatherings WHERE id = ?", id)

	var g gathering.Gathering
	err := row.Scan(&g.ID, &g.Title, &g.Latitude, &g.Longitude, &g.DateTime, &g.MaxParticipants, &g.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &g, nil
}

func (m *manager) UpdateGathering(gathering *gathering.Gathering) error {
	_, err := m.db.Exec("UPDATE gatherings SET title = ?, latitude = ?, longitude = ?, date_time = ?, max_participants = ? WHERE id = ?", gathering.Title, gathering.Latitude, gathering.Longitude, gathering.DateTime, gathering.MaxParticipants, gathering.ID)
	return err
}

func (m *manager) DeleteGathering(id int) error {
	_, err := m.db.Exec("DELETE FROM gatherings WHERE id = ?", id)
	return err
}

func (m *manager) CreateParticipant(participant *gathering.Participant) (int, error) {
	result, err := m.db.Exec("INSERT INTO gathering_participants (gathering_id, user_id) VALUES (?, ?)", participant.GatheringID, participant.UserID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *manager) ReadParticipant(id int) (*gathering.Participant, error) {
	row := m.db.QueryRow("SELECT * FROM gathering_participants WHERE id = ?", id)

	var p gathering.Participant
	err := row.Scan(&p.ID, &p.GatheringID, &p.UserID, &p.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (m *manager) UpdateParticipant(participant *gathering.Participant) error {
	_, err := m.db.Exec("UPDATE gathering_participants SET gathering_id = ?, user_id = ? WHERE id = ?", participant.GatheringID, participant.UserID, participant.ID)
	return err
}

func (m *manager) DeleteParticipant(id int) error {
	_, err := m.db.Exec("DELETE FROM gathering_participants WHERE id = ?", id)
	return err
}

func (m *manager) CountParticipantOnGathering(gatheringID int) (int, error) {
	row := m.db.QueryRow("SELECT COUNT(*) FROM gathering_participants WHERE gathering_id = ?", gatheringID)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
