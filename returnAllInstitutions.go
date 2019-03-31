package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Retrieve all Institutions
func returnAllInstitutions(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db, err := sql.Open("mysql", os.Getenv("ARCHIVESMAP_API_DB"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Select rows from database
	rows, err := db.Query("SELECT id, name, latitude, longitude, country FROM Institutions")
	if err != nil {
		log.Fatal(err)
	}

	var institutions []Institution
	for rows.Next() {
		err := rows.Scan(&id, &name, &latitude, &longitude, &country)
		if err != nil {
			log.Fatal(err)
		}

		// From rows to Institutions
		institutions = append(institutions, Institution{
			ID:        id,
			Name:      name,
			Latitude:  latitude,
			Longitude: longitude,
			Country:   country,
		})
	}

	// Return json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(institutions)
}
