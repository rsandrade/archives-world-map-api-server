package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Institution : the archives
type Institution struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Country   string `json:"country"`
	Distance  string `json:"distance"`
}

var id int
var name string
var latitude string
var longitude string
var country string

func main() {

	router := mux.NewRouter()
	//router.Host("archivesmap.org")
	//router.Schemes("https")

	//router.HandleFunc("/institution/{id}", returnInstitution)
	//router.HandleFunc("/institutions", returnAllInstitutions)
	router.HandleFunc("/nearestinstitutions/{longitude}/{latitude}", nearInstitutions)

	//router.ListenAndServe(":8443")
	http.ListenAndServeTLS(":8443", "server.crt", "server.key", router)
}
