// Package api implements the API endpoints needed for the restaurant reservation API.
package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*
	Patron Section
*/
type patron struct {
	Id int `json:"id"`
}

func getPatrons(c echo.Context) error {
	defer fmt.Println("Endpoint Hit: getPatrons")

	p := new(patron)

	if err := c.Bind(p); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}

func postPatrons(c echo.Context) error {
	defer fmt.Println("Endpoint Hit: postPatrons")

	p := new(patron)

	if err := c.Bind(p); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}

/*
	Estabish API with Route Handler
*/
func HandleRequests() {
	// create mux router
	router := echo.New()

	// middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// routes
	router.GET("/patrons", getPatrons)
	router.POST("/patrons", postPatrons)

	// start server
	router.Logger.Fatal(router.Start(":10000"))
}
