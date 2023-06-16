recreate:
	docker-compose up -d --force-recreate

container:
	docker exec -it api.go-api-rest.dev bash