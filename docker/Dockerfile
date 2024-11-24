# Etapa de build
FROM golang:1.23 AS builder

WORKDIR /app

# Instalar Protoc e Go Plugins
RUN apt-get update && apt-get install -y protobuf-compiler && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ENV PATH="$PATH:/go/bin"

# Baixar dependências do Go
COPY go.mod go.sum ./
RUN go mod download

# Copiar todo o código
COPY . ./

# Gerar arquivos gRPC
RUN protoc --go_out=. --go-grpc_out=. internal/delivery/grpc/order.proto

# Compilar a aplicação
RUN go build -o main cmd/main.go

# Etapa de execução
FROM debian:bookworm-slim AS app

WORKDIR /app

# Copiar o executável gerado na etapa anterior
COPY --from=builder /app/main .

# Expor portas necessárias para REST e gRPC
EXPOSE 8080 50051

# Comando para iniciar a aplicação
CMD ["./main"]
