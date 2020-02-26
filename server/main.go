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

	// router.HandleFunc("/occupancy/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Please specify a zone ID: %q", html.EscapeString(r.URL.Path))
	// })
	router.HandleFunc("/occupancy/", db.calcOccupancy)

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

type ErrorResponse struct {
	Response string `json: response`
}

func (db *DB) calcOccupancy(w http.ResponseWriter, r *http.Request) {
	log.Println("Calculating Occupancy")
	// vars := mux.Vars(r)
	vars := r.URL.Query()
	var (
		buildingID string
		floorID    string
		roomID     string
	)
	buildingIDs := vars["buildingID"]
	if len(buildingIDs) != 0 {
		buildingID = buildingIDs[0]
	} else {
		buildingID = ""
	}
	floorIDs := vars["floorID"]
	if len(floorIDs) != 0 {
		floorID = floorIDs[0]
	} else {
		floorID = ""
	}
	roomIDs := vars["roomID"]
	if len(roomIDs) != 0 {
		roomID = roomIDs[0]
	} else {
		roomID = ""
	}
	var q string
	if buildingID != "" {
		q = fmt.Sprintf(`from(bucket: "my-test-bucket")
  |> range(start: %s)
  |> filter(fn: (r) => r.BuildingID == "%s")
  |> last()`, "-5m", buildingID)
	} else if floorID != "" {
		q = fmt.Sprintf(`from(bucket: "my-test-bucket")
  |> range(start: %s)
  |> filter(fn: (r) => r.FloorID == "%s")
  |> last()`, "-5m", floorID)
	} else if roomID != "" {
		q = fmt.Sprintf(`from(bucket: "my-test-bucket")
  |> range(start: %s)
  |> filter(fn: (r) => r.RoomID == "%s")
  |> last()`, "-5m", roomID)
	} else {
		response := ErrorResponse{Response: "must provide a building, floor or room ID"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
		return
	}
	println(q)
	resp, err := db.Client.QueryCSV(context.Background(), q, "833c7fbc1d19c9be")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Executed query")
	var response Response
	readings := make([]SensorData, 0)
	// if !resp.Next() {
	// 	log.Fatal("ERROR: no data in database") // todo check if this is desired
	// } else {
	for resp.Next() {
		println("READING")
		reading := SensorData{}
		err = resp.Unmarshal(&reading)
		if err != nil {
			log.Fatal(err)
		}
		readings = append(readings, reading)
	}
	if len(readings) == 0 {
		log.Fatal("ERROR: no data in database") // todo check if this is desired
	} else if len(readings) > 1 {
		sumOcc := 0.0
		for _, each := range readings {
			sumOcc += each.Value
			response = Response{Occupancy: sumOcc}
		}
	} else {
		response = Response{Occupancy: readings[0].Value}
	}
	// }

	fmt.Printf("%v\n", readings)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
