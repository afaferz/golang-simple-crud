package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DbConfigs struct {
	User, Dbname, Password, Host, Sslmode string
}

func ConnectDB() *sql.DB {
	connection := DbConfigs{
		"postgres",
		"shop",
		"postgres",
		"localhost",
		"disable",
	}

	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", connection.User, connection.Dbname, connection.Password, connection.Host, connection.Sslmode)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	return db
}
