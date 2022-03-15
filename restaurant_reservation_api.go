package main

import (
	"fmt"

	"github.com/gthomas0/restaurant_reservation_api/api"
	"github.com/gthomas0/restaurant_reservation_api/database"
)

func main() {
	fmt.Println("Instantiating API")
	database.ConnectDB()
	database.MigrateDB()
	defer database.CloseDB()
	api.HandleRequests()
}
