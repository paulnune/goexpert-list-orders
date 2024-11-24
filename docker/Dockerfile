# Etapa de build
FROM golang:1.23 AS builder

WORKDIR /app

# Instalar Protoc e Go Plugins
RUN apt-get update && apt-get install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ENV PATH="$PATH:/go/bin"

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Gerar arquivos gRPC
RUN protoc --go_out=. --go-grpc_out=. internal/delivery/grpc/order.proto

# Build da aplicação
RUN go build -o main cmd/main.go

# Etapa de banco de dados
FROM postgres:14 AS db

# Copiar o script init.sql
COPY init.sql /docker-entrypoint-initdb.d/init.sql

# Configuração do banco
ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=orders

# Garantir permissões no init.sql
RUN chmod 644 /docker-entrypoint-initdb.d/init.sql

# Expor a porta padrão
EXPOSE 5432

# Etapa de aplicação
FROM debian:bookworm-slim AS app

WORKDIR /app

COPY --from=builder /app/main .
CMD ["./main"]
