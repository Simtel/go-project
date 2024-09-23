.DEFAULT_GOAL := help

help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

test:
	go test -v ./...

build:
	go build

run:
	go run .

watch:
	 ~/go/bin/gow run .

mocks:
	go generate ./...

up:
	docker compose up -d

down:
	docker compose down

cli: ## PHP console
	docker exec -it  go-project-go bash