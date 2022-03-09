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
type Patron struct {
	Id int `json:"id"`
}

func getPatrons(c echo.Context) error {
	defer fmt.Println("Endpoint Hit: getPatrons")

	patron := new(Patron)

	if err := c.Bind(patron); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, patron)
}

func postPatrons(c echo.Context) error {
	defer fmt.Println("Endpoint Hit: postPatrons")

	patron := new(Patron)

	if err := c.Bind(patron); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, patron)
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
