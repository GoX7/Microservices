package repository

import (
	"time"
)

type (
	Event struct {
		ID          string `json:"id"`
		Login       string `json:"login"`
		Event       string `json:"event"`
		Createstamp int64  `json:"createstamp"`
	}
)

func (pstgr *Postgres) CreateEvent(login string, event string) (int64, error) {
	var id int64
	err := pstgr.db.QueryRow("INSERT INTO events (login, event, createstamp) VALUES($1, $2, $3) RETURNING id", login, event, time.Now().Unix()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (pstgr *Postgres) GetEvents(login string) ([]*Event, error) {
	rows, err := pstgr.db.Query("SELECT * FROM events WHERE login=$1", login)
	if err != nil {
		return nil, err
	}

	var events []*Event
	for rows.Next() {
		var stamp int64
		var id, login, event string

		if err := rows.Scan(&id, &login, &event, &stamp); err != nil {
			return nil, err
		}

		events = append(events, &Event{
			ID: id, Createstamp: stamp,
			Login: login, Event: event,
		})
	}

	return events, nil
}

func (pstgr *Postgres) GetEvent(id string, login string) (*Event, error) {
	var event string
	var createstamp int64
	err := pstgr.db.QueryRow("SELECT event, createstamp FROM events WHERE id=$1 AND login=$2", id, login).Scan(
		&event, &createstamp,
	)

	if err != nil {
		return nil, err
	}

	return &Event{
		ID: id, Login: login,
		Event: event, Createstamp: createstamp,
	}, nil
}

func (pstgr *Postgres) DeleteEvent(id string, login string) error {
	_, err := pstgr.db.Exec("DELETE FROM events WHERE id=$1 AND login=$2", id, login)
	return err
}
