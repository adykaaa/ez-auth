include .env
export $(shell sed 's/=.*//' .env)

help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done
.PHONY: help

db: # Sets up PostgreSQL in a docker container
	docker run -e POSTGRES_USER=user -e POSTGRES_PASSWORD=pw -e POSTGRES_DB=notes -p 5432:5432 -d --name postgres-dev postgres
.PHONY: db

redis: # Sets up Redis in a docker container
	docker run -p 6379:6379 -d --name redis-dev redis
.PHONY: redis

sqlc: # Generates the DB backend code according to the sqlc.yaml file located in the root folder
	sqlc generate
.PHONY: sqlc

test: # Runs the unit tests
	go test -v -cover ./...
.PHONY: test

dbmock: # Generates the DB mocks
	mockgen -package mockdb -destination db/mock/querier.go  github.com/adykaaa/ez-auth/db/sqlc Querier
.PHONY: dbmock

run: # Builds and runs ez-auth inside a Docker container
	docker build . -t ezauth
	docker run ezauth -d -p $(HTTP_SERVER_ADDRESS):$(HTTP_SERVER_ADDRESS)
.PHONY: run

run-all: # Runs the docker compose file with all the services necessary for the application
	docker compose up -d
.PHONY: run-all