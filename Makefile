.DEFAULT_GOAL := help

help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

vtest:
	go test -v ./...

test:
	go test ./...

build:
	go build

run:
	go run . http

watch:
	 ~/go/bin/gow run . http

mocks:
	go generate ./...

up:
	docker compose up -d

down:
	docker compose down

cli: ## PHP console
	docker exec -it  go-project-go bash

lint:
	golangci-lint run

migrate:
	go run . migrate