package postgres

import (
	"database/sql"
	"financialManagement/internal/config"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct {
	db *sql.DB
}

func MustNewStorage(cfg config.PostgresConfig) *Storage {
	sourceName := cfg.GetDataSourceName()
	db := connectDB(sourceName)
	return &Storage{db}
}

func connectDB(sourceName string) *sql.DB {
	db, err := sql.Open("postgres", sourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Disconnect(storage *Storage) error {
	return storage.db.Close()
}
