package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/influxdata/influxdb-client-go"
)

// DB encapsulates the database connection so that it might be shared between response handlers
type DB struct {
	Client *influxdb.Client
}

//InfluxToken is injected at compile-time and authorises access to the Influx Database
var InfluxToken string

// NewDB creates a DB struct which abstracts the InfluxDB connection,
// allowing for http response methods to share the same connection
func NewDB(token string) DB {
	client := http.Client{}
	iDBClient, err := influxdb.New("https://us-central1-1.gcp.cloud2.influxdata.com", token, influxdb.WithHTTPClient(&client))
	if err != nil {
		log.Fatal(err)
	}
	return DB{Client: iDBClient}
}
func main() {
	db := NewDB(InfluxToken)
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You have reached the toplevel of the RoT-IoT internal web server, the following endpoints are available: %q", html.EscapeString(r.URL.Path))
	})

	router.HandleFunc("/occupancy/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Please specify a zone ID: %q", html.EscapeString(r.URL.Path))
	})
	router.HandleFunc("/occupancy/{zoneId}", db.calcOccupancy)

	log.Fatal(http.ListenAndServe(":6969", router))
}

// SensorData allows for important information from InfluxDB query responses to be accessed structurally
type SensorData struct {
	Time  time.Time `flux:"_time" json:"time"`
	Field string    `flux:"_field" json:"field"`
	Value float64   `flux:"_value" json:"value"`
}

// Response stores data to be returned to frontend via json encoding
type Response struct {
	Occupancy float64 `json:"occupancy"`
}

func (db *DB) calcOccupancy(w http.ResponseWriter, r *http.Request) {
	log.Println("Calculating Occupancy")
	vars := mux.Vars(r)
	zoneID := vars["zoneId"]
	q := fmt.Sprintf(`from(bucket: "my-test-bucket")
  |> range(start: %s)
  |> filter(fn: (r) => r.ID == "%s")
  |> last()`, "-5m", zoneID)
	resp, err := db.Client.QueryCSV(context.Background(), q, "833c7fbc1d19c9be")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Executed query")
	var response Response
	readings := SensorData{}
	if !resp.Next() {
		log.Fatal("ERROR: no data in database") // todo check if this is desired
	} else {
		err = resp.Unmarshal(&readings)
		if err != nil {
			log.Fatal(err)
		}
		response = Response{Occupancy: readings.Value}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
