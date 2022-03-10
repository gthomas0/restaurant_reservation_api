// Package models defines the tables in the database, and provides
// easy ways to connect to and manipulate the data.
package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

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
	Database Connection Section
*/
func ConnectDB() {
	host := "127.0.0.1"
	port := "5432"
	user := "postgres"
	password := "postgres"
	dbName := "restaurants"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	result, err := sql.Open("postgres", psqlInfo)
}
