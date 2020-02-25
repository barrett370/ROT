package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

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
	// we use client.NewRowMetric for the example because it's easy, but if you need extra performance
	// it is fine to manually build the []client.Metric{}.
	influx, err := DBConnect(INFLUX_TOKEN)
	if err != nil {
		panic(err)
	}
	source := rand.NewSource(10)
	rand := rand.New(source)
	spoof_data(influx, *rand, 100)
	// The actual write..., this method can be called concurrently.
}

func spoof_data(influx *influxdb.Client, r rand.Rand, repeats int) {

	report := SensorReport{
		CO2:         float64(r.Intn(101)) / 100,
		Temperature: float64(r.Intn(30-12+1) + 12),
	}
	fmt.Printf("%v\n", report)
	myMetric := []influxdb.Metric{
		influxdb.NewRowMetric(
			map[string]interface{}{"CO2": report.CO2, "temperature": report.Temperature},
			"Sensor Readings",
			map[string]string{"Hostname": "TestBox1"},
			time.Now()),
	}
	_, err := influx.Write(context.Background(), "my-test-bucket", "833c7fbc1d19c9be", myMetric...)
	if err != nil {
		log.Fatal(err) // as above use your own error handling here.
	}
	if repeats > 0 {
		//spoof_data(influx, r, repeats+-1)
		spoof_data(influx, r, repeats)
	}

	//

}
