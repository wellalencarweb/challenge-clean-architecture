.PHONY: up build run

up:
	docker-compose up -d

down:
	docker-compose down

build:
	go mod tidy

run:
	cd cmd/ordersystem && go run main.go wire_gen.go
	
test:
	go test ./... -v