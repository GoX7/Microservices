package repository

import (
	"database/sql"
	"fmt"
	"server/models"
	"time"

	_ "github.com/lib/pq"
)

type (
	Postgres struct {
		db *sql.DB
	}
)

func NewPostgres(config *models.Config, model *Postgres) {
	var connect *sql.DB
	var err error

	psc := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
		config.POSTGRES_HOST, config.POSTGRES_USER, config.POSTGRES_PASSWORD, config.POSTGRES_DBNAME)

	for i := 1; i <= 10; i++ {
		connect, err = sql.Open("postgres", psc)
		if err == nil {
			fmt.Println("[+] postgres.connect: ok")
			break
		}

		fmt.Printf("[-] postgres.connect: %d/10\n", i)
		time.Sleep(time.Second)
	}

	for i := 1; i <= 10; i++ {
		if err := connect.Ping(); err == nil {
			fmt.Println("[+] postgres.ping: ok")
			break
		}

		fmt.Printf("[-] postgres.ping: %d/10\n", i)
		time.Sleep(time.Second)
	}

	*model = Postgres{
		db: connect,
	}
}

func (pstgr *Postgres) SetTable() {
	pstgr.db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		login VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		createstamp INTEGER NOT NULL
	)`)
}
