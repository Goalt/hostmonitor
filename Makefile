run:
	docker-compose -p go-template -f docker-compose.yml up
clean:
	docker stop app ; docker rm app || true
	docker stop go-template_db_1; docker rm go-template_db_1 || true
	docker volume rm go-template_db_data || true
generate:
	cd internal/provider/ && wire ; cd ../..
lint:
	golangci-lint run
test:
	go test ./...
all: clean run
