.PHONY: build, run

go_build:
	go build -v ./cmd/web

build:
	docker-compose build

run:
	docker-compose up

down:
	docker-compose down

silent:
	docker-compose up -d

.DEFAULT_GOAL := run
