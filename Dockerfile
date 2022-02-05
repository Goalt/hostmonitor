FROM golang:1.17-alpine AS build_base

RUN apk add --no-cache git
RUN apk add --update make
RUN apk add --update alpine-sdk

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go get github.com/google/wire/cmd/wire

COPY . .

RUN make generate

# Build the Go apps
RUN go build -o app ./cmd/main.go

# Start fresh from a smaller image
FROM alpine:3.9

RUN apk add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /app
COPY --from=build_base /app ./

ENTRYPOINT ["/app/app"]