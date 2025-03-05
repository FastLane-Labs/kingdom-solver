include .env

run:
	OPERATIONS_RELAY_URL=$(OPERATIONS_RELAY_URL) go run main.go
