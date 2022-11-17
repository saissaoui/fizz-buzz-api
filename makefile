serve:
	go run main.go	
generate-mocks:
	go generate ./...	
tests:
	go test ./...	
make docker-dev:
	docker compose up -d redis
docker-run:
	docker compose up -d