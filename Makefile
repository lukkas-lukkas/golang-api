build-run:
	docker-compose up -d --build
	docker exec -it broker kafka-topics --create --topic golang-api_create-course --bootstrap-server localhost:9092 --replication-factor 1
	docker exec -i api.golang-api.dev go run main.go