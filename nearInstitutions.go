package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func nearInstitutions(w http.ResponseWriter, r *http.Request) {

	// Get the location from mobile using /nearestinstitutions?latitude=xxxxx&longitude=xxxxxx
	keys := mux.Vars(r)
	latmobile := keys["latitude"]
	longmobile := keys["longitude"]

	// Connect to database
	db, err := sql.Open("mysql", os.Getenv("ARCHIVESMAP_API_DB"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Compare the location sent by app with all institutions and get distance
	var institutions []Institution
	var meters string

	// SELECT ST_Distance_Sphere(point(longmobile, latmobile), point(longdb, latdb))
	// The result of SQL statement above is the distance in meters

	// Select all institutions and check distance one by one
	rows, _ := db.Query(
		"SELECT id, name, latitude, longitude, country, ST_Distance_Sphere("+
			"point(?, ?), point(longitude, latitude)) AS meters "+
			"FROM Institutions "+
			"ORDER BY ST_Distance_Sphere("+
			"point(?, ?), point(longitude, latitude)) "+
			"LIMIT 10", longmobile, latmobile, longmobile, latmobile)

	for rows.Next() {
		err := rows.Scan(&id, &name, &latitude, &longitude, &country, &meters)
		if err != nil {
			log.Fatal(err)
		}

		// From row to Institution
		institutions = append(institutions, Institution{
			ID:        id,
			Name:      name,
			Latitude:  latitude,
			Longitude: longitude,
			Country:   country,
			Distance:  meters,
		})
	}

	// Send back to app the nearest institutions in json format
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(institutions)

}
