.PHONY: go_build
go_build:
	go build -v ./cmd/web

.PHONY: build
build:
	docker-compose build

.PHONY: run
run:
	docker-compose up

.PHONY: down
down:
	docker-compose down

.PHONY: silent
silent:
	docker-compose up -d

.DEFAULT_GOAL := run
