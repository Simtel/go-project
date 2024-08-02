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