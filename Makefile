include .env

.PHONY: build-server
build-server:
	@cd server && go build -ldflags "-X main.InfluxToken=${INFLUX_READ_TOKEN}" -o server && mv ./server ../bin
	@echo "\033[0;32m SUCCESS Compiled server binary \033[0m"

.PHONY: server
server: 
	@make build-server 
	@echo "\033[0;32m Running server \033[0m"
	@cd ./bin && ./server

.PHONY: deps
deps:
	@echo "\033[1;33m Installing dependencies \033[0m"
	cd ./server && go mod download 
	cd ./web/rot && sudo npm i 

.PHONY: pretty
pretty: 
	@goimports -w ./
	@gofmt -w ./
	@echo "\033[0;31m ------------------------------------------------\033[0m"
	@echo "\033[0;35m REMINDER: Check go import blocks for blank lines\033[0m"
	@echo "\033[0;31m ------------------------------------------------\033[0m"
