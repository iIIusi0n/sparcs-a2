package data

import (
	"api-server/controllers/gathering"
)

func (m *manager) CreateGathering(gathering *gathering.Gathering) (int, error) {
	result, err := m.db.Exec("INSERT INTO gatherings (user_id, title, latitude, longitude, date_time, duration_in_hours, max_participants, thumbnail_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", gathering.UserID, gathering.Title, gathering.Latitude, gathering.Longitude, gathering.DateTime, gathering.DurationInHours, gathering.MaxParticipants, gathering.ThumbnailURL)
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

func (m *manager) ReadGatherings() ([]*gathering.Gathering, error) {
	rows, err := m.db.Query("SELECT * FROM gatherings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	gatherings := make([]*gathering.Gathering, 0)
	for rows.Next() {
		var g gathering.Gathering
		err := rows.Scan(&g.ID, &g.Title, &g.Latitude, &g.Longitude, &g.DateTime, &g.MaxParticipants, &g.CreatedAt)
		if err != nil {
			return nil, err
		}
		gatherings = append(gatherings, &g)
	}

	return gatherings, nil
}

func (m *manager) UpdateGathering(gathering *gathering.Gathering) error {
	_, err := m.db.Exec("UPDATE gatherings SET title = ?, latitude = ?, longitude = ?, date_time = ?, max_participants = ? WHERE id = ?", gathering.Title, gathering.Latitude, gathering.Longitude, gathering.DateTime, gathering.MaxParticipants, gathering.ID)
	return err
}

func (m *manager) DeleteGathering(id int) error {
	_, err := m.db.Exec("DELETE FROM gatherings WHERE id = ?", id)
	return err
}

func (m *manager) CountGatheringOnUser(userID int) (int, error) {
	row := m.db.QueryRow("SELECT COUNT(*) FROM gatherings WHERE user_id = ?", userID)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
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

func (m *manager) CheckUserParticipatedGathering(userID, gatheringID int) (bool, error) {
	row := m.db.QueryRow("SELECT COUNT(*) FROM gathering_participants WHERE user_id = ? AND gathering_id = ?", userID, gatheringID)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (m *manager) CreateGatheringLocation(location *gathering.Location) (int, error) {
	result, err := m.db.Exec("INSERT INTO gathering_locations (gathering_id, post_id) VALUES (?, ?)", location.GatheringID, location.PostID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *manager) ReadGatheringLocation(id int) (*gathering.Location, error) {
	row := m.db.QueryRow("SELECT * FROM gathering_locations WHERE id = ?", id)

	var l gathering.Location
	err := row.Scan(&l.ID, &l.GatheringID, &l.PostID, &l.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &l, nil
}

func (m *manager) UpdateGatheringLocation(location *gathering.Location) error {
	_, err := m.db.Exec("UPDATE gathering_locations SET gathering_id = ?, post_id = ? WHERE id = ?", location.GatheringID, location.PostID, location.ID)
	return err
}

func (m *manager) DeleteGatheringLocation(id int) error {
	_, err := m.db.Exec("DELETE FROM gathering_locations WHERE id = ?", id)
	return err
}

func (m *manager) ReadGatheringLocationByGatheringID(gatheringID int) (*gathering.Location, error) {
	row := m.db.QueryRow("SELECT * FROM gathering_locations WHERE gathering_id = ?", gatheringID)

	var l gathering.Location
	err := row.Scan(&l.ID, &l.GatheringID, &l.PostID, &l.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &l, nil
}
