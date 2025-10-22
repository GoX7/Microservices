package repository

import "time"

func (pstgr *Postgres) RegisterUser(login string, password string) (int64, error) {
	var id int64
	err := pstgr.db.QueryRow("INSERT INTO users (login, password, createstamp) VALUES($1, $2, $3) RETURNING id",
		login, password, time.Now().Unix()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (pstgr *Postgres) LoginUser(login string, password string) (int64, error) {
	var id int64
	err := pstgr.db.QueryRow("SELECT id FROM users WHERE login=$1 AND password=$2",
		login, password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
