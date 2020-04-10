package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/influxdata/influxdb-client-go"
)

type numberSet struct {
	values []float64
	bounds []float64
}

func newNumberSet(values []float64, weight []float64) (*numberSet, error) {
	if len(values) != len(weight) {
		return nil, fmt.Errorf("values and weight should have the same length")
	}
	s := &numberSet{
		values: values,
		bounds: weight,
	}
	sort.Sort(s)

	sum := float64(0)
	for i, weight := range s.bounds {
		sum += weight
		s.bounds[i] = sum
	}
	if sum-1 > 1e9 {
		return nil, fmt.Errorf("sum of weight should be 1, but was %f", sum)
	}
	return s, nil
}

func (s *numberSet) Len() int { return len(s.values) }
func (s *numberSet) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
	s.bounds[i], s.bounds[j] = s.bounds[j], s.bounds[i]
}
func (s *numberSet) Less(i, j int) bool { return s.bounds[i] < s.bounds[j] }

// Generator is a struct that can returns a random number chosen from a set
// of numbers where each has a specified probability.
type Generator struct {
	randSource rand.Source
	size       int
	numberSet
}

// NewGenerator return a Generator. It returns an error if len(weight) != len(values),
// or if the sum of weights is != 1.
// Two Generators with same seed, values and weight will always produce the same sequence
// of random number
func NewGenerator(seed int64, values []float64, weight []float64) (*Generator, error) {
	s, err := newNumberSet(values, weight)
	if err != nil {
		return nil, err
	}
	return &Generator{
		randSource: rand.NewSource(seed),
		size:       len(values),
		numberSet:  *s,
	}, nil
}

// Random returns a random number from the generator number set.
func (g *Generator) Random() float64 {
	r := float64(g.randSource.Int63()) / (1 << 63)
	i := sort.Search(g.size, func(i int) bool {
		return g.bounds[i] >= r
	})
	return g.values[i]
}

var INFLUX_TOKEN string
var buildingID int
var floorID int
var roomID int

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
	flag.IntVar(&buildingID, "b", -1, "Building ID")
	flag.IntVar(&floorID, "f", -1, "Floor ID")
	flag.IntVar(&roomID, "r", -1, "Room ID")
	var val int
	flag.IntVar(&val, "v", 0, "Val to set")

	flag.Parse()
	if buildingID == -1 || floorID == -1 || roomID == -1 {
		log.Fatal("Please enter valid building, floor and room ids")
	}

	influx, err := DBConnect(INFLUX_TOKEN)
	if err != nil {
		panic(err)
	}
	//source := rand.NewSource(10)
	// rand := rand.New(source)
	if val == 0 {
		rand, err := NewGenerator(10, []float64{-1.0, 0.0, 1.0}, []float64{0.25, 0.5, 0.25})
		if err != nil {
			log.Fatal(err)
		}
		spoof_data(influx, *rand, 100, 10)
	} else {
		rand, err := NewGenerator(10, []float64{float64(val)}, []float64{1})
		if err != nil {
			log.Fatal(err)
		}
		spoof_data(influx, *rand, 0, 0)
	}

	// The actual write..., this method can be called concurrently.
}

func spoof_data(influx *influxdb.Client, r Generator, repeats int, prev_occ float64) {

	//report := SensorReport{
	//	CO2:         float64(r.Intn(101)) / 100,
	//	Temperature: float64(r.Intn(30-12+1) + 12),
	//}

	//fmt.Printf("%v\n", report)
	fmt.Printf("BID: %s, FID: %s, RID: %s\n", strconv.Itoa(buildingID), strconv.Itoa(floorID), strconv.Itoa(roomID))
	occupancy := prev_occ + r.Random()
	fmt.Printf("%f", occupancy)
	myMetric := []influxdb.Metric{
		influxdb.NewRowMetric(
			map[string]interface{}{"occupancy": occupancy},
			"Sensor Readings",
			map[string]string{"Hostname": "TestBox1", "BuildingID": strconv.Itoa(buildingID), "FloorID": strconv.Itoa(floorID), "RoomID": strconv.Itoa(roomID)},
			time.Now()),
	}
	// _, err := influx.Write(context.Background(), "sjb786's Bucket", "833c7fbc1d19c9be", myMetric...)
	_, err := influx.Write(context.Background(), "my-test-bucket", "833c7fbc1d19c9be", myMetric...)
	if err != nil {
		log.Fatal(err) // as above use your own error handling here.
	}
	if repeats > 0 {
		//spoof_data(influx, r, repeats+-1)
		spoof_data(influx, r, repeats+-1, occupancy)
	}

	//

}
