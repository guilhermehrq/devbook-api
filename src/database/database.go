package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// Connect opens a connection with the database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
