FROM mcr.microsoft.com/vscode/devcontainers/go:1.17

RUN apt update
RUN apt install -y protobuf-compiler libprotobuf-dev

RUN PREFIX="/usr/local" && \
    VERSION="1.0.0-rc12" && \
    curl -sSL \
        "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m).tar.gz" | \
        tar -xvzf - -C "${PREFIX}" --strip-components 1

USER vscode

WORKDIR /workspace

RUN go get github.com/cweill/gotests/gotests
RUN go get github.com/google/wire/cmd/wire
RUN go get github.com/golang/mock/gomock
RUN go get github.com/golang/mock/mockgen

COPY . .

RUN go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc