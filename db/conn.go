package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	Host     = "go_db"
	Port     = 5432
	User     = "postgres"
	Password = "1234"
	DbName   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, DbName)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to " + DbName)

	return db, nil
}
