phony: run test

run:
	go run cmd/telegram_service/main.go

test:
	go test ./...