SHELL := /bin/bash

copy-config:
	cp configs/env.local.template .env

dev:
	@echo "Going to run go mod tidy verify"
	@go mod tidy
	@go mod verify
	@go mod vendor
	swag init -g cmd/api/main.go

run-dev:
	swag init -g cmd/api/main.go
	go run cmd/api/main.go
