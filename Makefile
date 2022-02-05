run:
	docker-compose -p hostmonitor -f docker-compose.yml up
clean:
	docker stop app ; docker rm app || true
lint:
	golangci-lint run
test:
	go test ./...
all: clean run
