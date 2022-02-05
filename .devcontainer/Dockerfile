FROM mcr.microsoft.com/vscode/devcontainers/go:1.17

RUN apt update
RUN apt install -y protobuf-compiler

USER vscode

WORKDIR /workspace

RUN go get github.com/cweill/gotests/gotests
RUN go get github.com/google/wire/cmd/wire
RUN go get github.com/golang/mock/gomock
RUN go get github.com/golang/mock/mockgen

COPY . .

# RUN go mod tidy

# RUN go install \
#     github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
#     github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
#     github.com/golang/protobuf/protoc-gen-go

EXPOSE 8080