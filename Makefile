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
	# mkdir -p "api_pb"
	# protoc -I/usr/local/include -I. \
    #         -I${GOPATH}/src \
    #         -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    #         -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
    #         --grpc-gateway_out=logtostderr=true:./api_pb \
    #         --swagger_out=allow_merge=true,merge_file_name=api:. \
    #         --go_out=plugins=grpc:./api_pb ./internal/infrastructure/grpc/api.proto

all: clean run
