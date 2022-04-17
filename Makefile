.PHONY: run tests

OS_NAME := $(shell uname)

ifeq ($(OS_NAME), Darwin)
	DB_HOST := "docker.for.mac.localhost"
else
	DB_HOST := "localhost"
endif

run:
	@go run cmd/main.go

tests:
	@go test -v ./...
