build:
	go build -o bin/main.out ./cmd/main.go

run:
	./bin/main.out

clear:
	rm -rf ./bin

redis_dev:
	docker run --name redis-go-dev -p 6379:6379 -d redis