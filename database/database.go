package database

import (
	"database/sql"
	"fmt"
	"time"
)

func GetConnection() *sql.DB {
	dns := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		"172.18.128.131", "5432", "postgres", "postgres", "go_database", "disable",
	)

	db, err := sql.Open("pgx", dns)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
