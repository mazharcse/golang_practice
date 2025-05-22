package utils

import (
	"database/sql"
	"log/slog"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func createDB() *DB {
	dbAddress := os.Getenv("DB_ADDRESS")
	if dbAddress == "" {
		panic("invalid database address")
	}
	slog.Info("Connecting to database " + dbAddress)
	db, err := sql.Open("postgres", dbAddress)

	if err != nil {
		slog.Error("Fail to connect db " + err.Error())
		return nil
	}

	slog.Info("Successfully connected to db")
	return &DB{db}
}

var dbOnce sync.Once
var appDB *DB

func GetDB() *DB {
	dbOnce.Do(func ()  {
		appDB = createDB()
	})

	return appDB
}