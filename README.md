# RoT: Room Occupancy Tracking

## Setup

To build and run the go web server:

1. cd into `./server` and run `go mod download` to install dependencies.
2. return to the toplevel dir
3. create a file `.env` and add `INFLUX_READ_TOKEN` and an approprite token to the file
4. run `make build-server` to compile a binary to `./bin` or `make server` to build and run the binary
