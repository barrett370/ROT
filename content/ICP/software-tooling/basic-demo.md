---
title: Basic InfluxDB implementation
---

After our initial research into InfluxDB and it's client libraries, we decided to complete a trial implementation using Golang as it offered a fast, compiled language that lent itself well to embedded systems and has excellent networking and json parsing libraries.

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	fsnotify "github.com/fsnotify/fsnotify"
	influxdb "github.com/influxdata/influxdb-client-go"
)

var INFLUX_TOKEN string

func DBConnect() (*influxdb.Client, error) {
	client := http.Client{}
	return influxdb.New("https://us-central1-1.gcp.cloud2.influxdata.com", INFLUX_TOKEN, influxdb.WithHTTPClient(&client))
}

type SensorReport struct {
	CO2         float64 `json:"CO2"`
	Temperature float64 `json:"Temperature"`
}

func main() {
	DBConnect()
	influx, err := DBConnect()
	if err != nil {
		panic(err)
	}
	watch_sensors(influx)
}

func watch_sensors(influx *influxdb.Client) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Printf("EVENT! %#v\n", event)
				readings, err := os.Open("sensor-readings.json")
				if err != nil {
					log.Fatal(err)
				}
				defer readings.Close()
				var report SensorReport
				fileBytes, err := ioutil.ReadAll(readings)
				if err != nil {
					log.Fatal(err)
				}
				json.Unmarshal(fileBytes, &report)
				fmt.Printf("%+v\n", report)
				myMetric := []influxdb.Metric{
					influxdb.NewRowMetric(
						map[string]interface{}{"CO2": report.CO2, "temperature": report.Temperature},
						"Sensor Readings",
						map[string]string{"Hostname": "TestBox1"},
						time.Now()),
				}
				_, err = influx.Write(context.Background(), "my-test-bucket", "833c7fbc1d19c9be", myMetric...)
				if err != nil {
					log.Fatal(err) 
				}
				watch_sensors(influx)
				
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add("sensor-readings.json"); err != nil {
		fmt.Println("ERROR", err)
	}

	<-done
}
```

### Contents of `sensor-readings.json`

```json 

{
	"CO2": 0.2,
	"Temperature": 23
}
```

The code above watches for changes in a given file `sensor-readings.json` and upon changes unmarshals the contents to a Go struct. This struct is then used to fill out a influxDB metric for writing to a given bucket. 
This code is a proof of concept, showing that if we were to record sensor data to a json structured format, we could easily and reliably push this data to our database retaining its structure and time sensitive nature. Running this script results in the following results being displayed from within InfluxDBs query dashboard:

![InfluxDB Dashboard](/images/uploads/IDBDashboard.png)

Another advantage of a compiled language such as Go, is that we can make use of *compile time variable injection* to embed sensitive information such as our database access token into the binary which can then be placed on our embedded system without the need for a plaintext version of the token to be stored on these devices in the public domain. Go also produces a statically linked compiled binary allowing for it to be run on any supporting system without additional software or libraries to be installed along with it.