build:
	go build -o bin/main.out main.go

run:
	./bin/main.out

redis_dev:
	docker run --name redis-go-dev -p 6379:6379 -d redis