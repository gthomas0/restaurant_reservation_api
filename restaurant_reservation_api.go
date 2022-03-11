package main

import (
	"fmt"

	"github.com/gthomas0/restaurant_reservation_api/api"
)

func main() {
	fmt.Println("Instantiating API")
	api.HandleRequests()
}
