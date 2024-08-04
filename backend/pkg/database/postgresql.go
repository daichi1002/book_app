package database

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgreAQL(dsn string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Connection retry logic
	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(time.Second * 3)
	}

	if err != nil {
		return nil, err
	}
	return db, nil
}
