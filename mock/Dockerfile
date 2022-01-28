FROM golang:1.17.6-alpine3.15 AS builder
WORKDIR /build
COPY main.go main.go
COPY data.go data.go
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
RUN go build -o server

FROM alpine:3.15
WORKDIR /app
COPY --from=builder build/server server
CMD ["./server"]
