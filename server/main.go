package main

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/influxdata/influxdb-client-go"
)

type DB struct {
	Client *influxdb.Client
}

var INFLUX_TOKEN string

func NewDB(token string) DB {
	client := http.Client{}
	iDBClient, err := influxdb.New("https://us-central1-1.gcp.cloud2.influxdata.com", token, influxdb.WithHTTPClient(&client))
	if err != nil {
		log.Fatal(err)
	}
	return DB{Client: iDBClient}
}
func main() {
	db := NewDB(INFLUX_TOKEN)
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

type SensorData struct {
	Time  time.Time `flux:"_time" json:"time"`
	Field string    `flux:"_field" json:"field"`
	Value float64   `flux:"_value" json:"value"`
}

func (db *DB) calcOccupancy(w http.ResponseWriter, r *http.Request) {
	log.Println("Calculating Occupancy")
	vars := mux.Vars(r)
	zoneID := vars["zoneId"]
	q := fmt.Sprintf(`from(bucket: "my-test-bucket")
  |> range(start: %s)
  |> filter(fn: (r) => r.ID == "%s")
  |> last()`, "-5m", zoneID)
	log.Printf("Execting query %s\n", q)
	resp, err := db.Client.QueryCSV(context.Background(), q, "833c7fbc1d19c9be")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Executed query")
	// readings := make([]SensorData, 0)
	var occupancies []float64
	readings := SensorData{}
	for resp.Next() {
		err = resp.Unmarshal(&readings)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%v\n", readings)
		occupancies = append(occupancies, readings.Value)
	}
	fmt.Printf("%v", occupancies)
	avg := func(xs []float64) float64 {
		sum := 0.0
		for _, each := range xs {
			sum += each
		}
		log.Printf("total sum: %f", sum)
		return sum / float64(len(xs))
	}
	log.Printf("%f\n", avg(occupancies))
	fmt.Fprintf(w, `{"occupancy" : %f}`, occupancies[0])
}
