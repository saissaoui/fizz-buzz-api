serve:
	go run main.go	
test:
	go test ./...	
make docker-dev:
	docker compose up -d redis
docker-run:
	docker compose up -d