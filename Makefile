include .env
build-server:
	cd server && go build -ldflags "-X main.InfluxToken=${INFLUX_READ_TOKEN}" && mv ./server ../bin

.PHONY: server
server: 
	make build-server && cd ./bin && ./server

.PHONY: deps
deps:
	cd ./server && go mod download 
	cd ./web/rot && sudo npm i 
