package config

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"os"
)

var database *sql.DB

func InitDB() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, pass, host, port, name)

	db, err := sql.Open("sqlserver", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	database = db

	log.Print("Database connection established!")
}

func DestroyDB(db *sql.DB) {
	err := db.Close()

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Database connection destroyed!")
}

func GetDBConnection() *sql.DB {
	return database
}
