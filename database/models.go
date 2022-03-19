// Package database defines the tables in the database, and provides
// easy ways to connect to and manipulate the data.
package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

/*
	Database Variable
*/
var db *sql.DB

/*
	Model Section
*/
type Patron struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Number string `json:"number"`
	Email  string `json:"email"`
}

type Restaurant struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Sunday    string `json:"sunday"`
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Saturday  string `json:"saturday"`
	Hours     string `json:"hours"`
}

/*
	Table Function Section
*/
func InsertPatron(class Patron) {

}

/*
	Database Connection Section
*/
func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, name)
	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("The database failed to connect: %s", err)
	}

	db = database
}

func CloseDB() error {
	return db.Close()
}

func MigrateDB() {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to grab the database driver: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///restaurant_reservation_api/database/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Failed to initialize the migration handler: %s", err)
	}

	m.Up()
}
