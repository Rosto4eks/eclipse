build:
	go build -o ./bin/server -v ./cmd/server

run:
	sudo ./bin/server

br:
	go build -o ./bin/server -v ./cmd/server
	sudo ./bin/server

.DEFAULT_GOAL := br

