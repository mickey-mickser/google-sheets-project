.PHONY: build run build-container test
.SILENT:

build:
	go build -o ./.bin/service cmd/main.go

run: build
	./.bin/service

build-container:
	sudo docker build -t google-sheets-project .

test:
	go test ./...