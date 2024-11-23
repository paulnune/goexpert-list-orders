FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o main cmd/main.go

FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080 50051
CMD ["./main"]
