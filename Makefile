.PHONY: lint test

run:
	docker-compose -p hostmonitor -f docker-compose.yml up
clean:
	docker stop app ; docker rm app || true
lint:
	golangci-lint run
test:
	go test ./...
generate:
	cd internal/provider/ && wire ; cd ../..

	cd internal/infrastructure/grpc && buf mod update && buf generate ; cd ../../..

	mockgen -source=internal/usecase/interactor/statistics.go -destination=internal/mocks/interactor/statistics.go
all: clean run
