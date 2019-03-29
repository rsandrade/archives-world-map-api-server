package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Retrieve a Institution
func returnInstitution(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Connect to database
	db, err := sql.Open("mysql", os.Getenv("ARCHIVESMAP_API_DB"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var institution []Institution
	rows, _ := db.Query("SELECT id, name, latitude, longitude, country FROM Institutions WHERE id = ?", vars["id"])
	for rows.Next() {
		err := rows.Scan(&id, &name, &latitude, &longitude, &country)
		if err != nil {
			log.Fatal(err)
		}

		// From row to Institution
		institution = append(institution, Institution{ID: id, Name: name, Latitude: latitude, Longitude: longitude, Country: country})
	}

	// Return json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(institution)
}
