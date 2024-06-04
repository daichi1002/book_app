package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func ConnectDB() (*sql.DB, error) {
	// データベースへの接続情報
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		return nil, err
	}

	return db, nil
}
