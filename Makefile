build:
	go build -o bin/main.out ./cmd/main.go

run:
	./bin/main.out

clear:
	rm -rf ./bin

dev:
	rm -rf ./bin && go build -o bin/main.out ./cmd/main.go && ./bin/main.out 

test:
	go test ./...

redis_dev:
	docker run --name redis-go-dev -p 6379:6379 -d redis

postgres_dev:
	docker run -d \
			-p5432:5432 \
			-e POSTGRES_DB=db \
			-e POSTGRES_USER=postgres \
			-e POSTGRES_PASSWORD=postgres \
			postgres:latest