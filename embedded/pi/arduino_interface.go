package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/huin/goserial"
	"github.com/influxdata/influxdb-client-go"
)

func findArduino() string {
	log.Print("Searching for arduino")
	contents, err := ioutil.ReadDir("/dev/")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range contents {
		if strings.Contains(f.Name(), "tty.usbserial") || strings.Contains(f.Name(), "ttyUSB") {
			log.Print(f.Name())
			return "/dev/" + f.Name()
		}
	}
	return ""
}
func readSerial(port io.ReadWriteCloser) {
	buf := make([]byte, 128)
	n, _ := port.Read(buf)
	if n != 0 {
		log.Printf("%q\n", buf[:n])
	}
	readSerial(port)
}

// SensorData is a Struct for storing information resulting from InfluxDB Query
type SensorData struct {
	Time       time.Time `flux:"_time" json:"time"`
	Field      string    `flux:"_field" json:"field"`
	Value      float64   `flux:"_value" json:"value"`
	BuildingID int       `flux:"BuildingID" json:"buildingID"`
	FloorID    int       `flux:"FloorID" json:"floorID"`
	RoomID     int       `flux:"RoomID" json:"roomID"`
}

// DBConnect given an access token creates an influxdb client for interfacing with the DB
func DBConnect(InfluxToken string) (*influxdb.Client, error) {
	// You can generate a Token from the "Tokens Tab" in the UI
	client := http.Client{}
	return influxdb.New("https://us-central1-1.gcp.cloud2.influxdata.com", InfluxToken, influxdb.WithHTTPClient(&client))
}
func initRoomCounter(db *influxdb.Client) (int, error) {
	println("initialising room counter")
	q := fmt.Sprintf(`from(bucket: "my-test-bucket")
  |> range(start: %s)
  |> filter(fn: (r) => r.RoomID == "%s")
  |> last()`, "-30d", RoomID)

	resp, err := db.QueryCSV(context.Background(), q, "833c7fbc1d19c9be")
	if err != nil {
		return 0, err
	}
	readings := make([]SensorData, 0)
	for resp.Next() {
		reading := SensorData{}
		err = resp.Unmarshal(&reading)
		if err != nil {
			return 0, err
		}
		readings = append(readings, reading)
	}
	if len(readings) == 0 {
		log.Println("No preexisting data for room")
		return 0, nil
	} else if len(readings) > 1 {
		log.Println("multiple rooms sharing and ID")
		BuildingID = strconv.Itoa(readings[0].BuildingID)
		fmt.Printf("Initial Occupancy: %f", readings[0].Value)
		FloorID = strconv.Itoa(readings[0].FloorID)
		if readings[0].Value > 1000{
			return 0,nil
		}
		return int(readings[0].Value), nil
	} else {
		BuildingID = strconv.Itoa(readings[0].BuildingID)
		FloorID = strconv.Itoa(readings[0].FloorID)
		fmt.Printf("Initial Occupancy: %f", readings[0].Value)
		if readings[0].Value > 1000{
			return 0,nil
		}
		return int(readings[0].Value), nil
	}
}

func updateDBRoomCounter(db *influxdb.Client, occupancy int) error {
	myMetric := []influxdb.Metric{
		influxdb.NewRowMetric(
			map[string]interface{}{"occupancy": occupancy},
			"Sensor Readings",
			map[string]string{"BuildingID": BuildingID, "FloorID": FloorID, "RoomID": RoomID},
			time.Now()),
	}
	_, err := db.Write(context.Background(), "my-test-bucket", "833c7fbc1d19c9be", myMetric...)
	return err
}

var (
	// InfluxToken allows access to the InfluxDB database BUILDTIME INJECTION
	InfluxToken string
	// RoomID stores the roomID of the room the system running this script is deployed to BUILDTIME INJECTION
	RoomID string
	// FloorID stores the floorID of the room specified by RoomID
	FloorID string
	// BuildingID stores the buildingID of the room specified by RoomID
	BuildingID string
)

func main() {
	fmt.Printf("Setup for room: %s\n", RoomID)
	println(InfluxToken)
	db, err := DBConnect(InfluxToken)
	if err != nil {
		log.Print(err)
	}
	roomCounter, err := initRoomCounter(db)
	fmt.Printf("Setup for Building: %s, floor: %s\n", BuildingID, FloorID)
	if err != nil {
		log.Print(err)
	}
	c := &goserial.Config{Name: findArduino(), Baud: 9600}
	s, err := goserial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	for {

		buf := make([]byte, 128)
		n, _ := s.Read(buf)
		sval := string(buf[:n])
		ival, err := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(sval, "\n", ""), "\r", ""))

		if err != nil {
			log.Print(err)
		}
		log.Println(ival)
		if !(ival > 1) {
			roomCounter += ival
		}
		err = updateDBRoomCounter(db, roomCounter)
		if err != nil {
			log.Print(err)
		}

	}
}
