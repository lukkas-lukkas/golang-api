build:
	docker-compose up -d --build

up:
	docker-compose up -d --build

run:
	docker exec -it api.go-api-rest.dev go run main.go

tests:
	docker exec -it api.go-api-rest.dev go test ./...