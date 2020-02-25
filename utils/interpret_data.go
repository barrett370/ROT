package main

import (
	"net/http"

	"github.com/influxdata/influxdb-client-go"
)

var INFLUX_TOKEN string

type SensorReport struct {
	CO2         float64 `json:"CO2"`
	Temperature float64 `json:"Temperature"`
}

func DBConnect(InfluxToken string) (*influxdb.Client, error) {
	// You can generate a Token from the "Tokens Tab" in the UI
	client := http.Client{}
	return influxdb.New("https://us-central1-1.gcp.cloud2.influxdata.com", InfluxToken, influxdb.WithHTTPClient(&client))
}

func main() {

}
