package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
	Patron Section
*/
func getPatrons(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	defer fmt.Println("Endpoint Hit: getPatrons")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(body)
}

func postPatrons(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	defer fmt.Println("Endpoint Hit: postPatrons")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(body)
}

/*
	Estabish API with Route Handler
*/
func HandleRequests() {
	// create mux router
	router := mux.NewRouter()

	// handle endpoints
	router.HandleFunc("/patrons", getPatrons).Methods("GET")
	router.HandleFunc("/patrons", postPatrons).Methods("POST")

	// set handler to use /
	http.Handle("/", router)

	// log fatal errors
	log.Fatal(http.ListenAndServe(":10000", nil))
}
