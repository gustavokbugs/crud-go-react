package sqlc

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type database struct {
	DB *sql.DB
}

func NewDB() *database {
	return &database{}
}

func (d *database) Connect() (*sql.DB, error) {

	var err error

	d.DB, err = sql.Open("postgres", "postgresql://postgres:postgres@postgres:5432/compartilha-tech?sslmode=disable")
	if err != nil {
		log.Fatalln("db connect: ", err)
		return nil, err
	}

	err = d.DB.Ping()
	if err != nil {
		log.Fatalln("db ping: ", err)
		return nil, err
	}

	return d.DB, nil
}
