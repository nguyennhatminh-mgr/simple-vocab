FROM golang:alpine3.18 AS builder

WORKDIR /build

ENV PROJECT_DIR=/build
ENV GO111MODULE=on
ENV CGO_ENABLED=0

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
COPY .env .

#Setup hot-reload for dev stage
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

# Command to run when starting the container.
ENTRYPOINT CompileDaemon --polling --build="go build -o apiserver ." --command=./apiserver