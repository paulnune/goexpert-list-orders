FROM golang:1.23 AS builder
WORKDIR /app

# Instala o protoc e seus plugins
RUN apt-get update && apt-get install -y protobuf-compiler

# Instala os plugins para Go e gRPC
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Adiciona o diretório GOPATH/bin ao PATH
ENV PATH="$PATH:/go/bin"

# Copia os arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . ./

# Gera os arquivos gRPC
RUN protoc --go_out=. --go-grpc_out=. internal/delivery/grpc/order.proto

# Compila a aplicação
RUN go build -o main cmd/main.go

# Imagem final minimalista
FROM gcr.io/distroless/base
WORKDIR /
COPY --from=builder /app/main .
CMD ["./main"]
