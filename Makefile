SHELL := /bin/bash

# Load environment variables from .env file
include .env
export

.PHONY: migrate

migrate-up:
	migrate -database "postgres://${DATASOURCE_USERNAME}:${DATASOURCE_PASSWORD}@${DATASOURCE_HOST}:${DATASOURCE_PORT}/${DATASOURCE_DB_NAME}?sslmode=disable" -path db/migrations up

migrate-down:
	migrate -database "postgres://${DATASOURCE_USERNAME}:${DATASOURCE_PASSWORD}@${DATASOURCE_HOST}:${DATASOURCE_PORT}/${DATASOURCE_DB_NAME}?sslmode=disable" -path db/migrations down
