
# 🚀 Desafio GoExpert - Listagem de Orders

Este projeto foi desenvolvido por **Paulo Nunes**.

---

## 📝 Sobre o Desafio

Para este desafio, é necessário criar o usecase de listagem das orders. Esta listagem precisa ser feita com:

- 📡 **Endpoint REST** (GET `/order`)
- 🔗 **Service ListOrders com GRPC**
- 🧩 **Query ListOrders GraphQL**

Criar as migrações necessárias e o arquivo `api.http` com as requests para criar e listar as orders.

Para a criação do banco de dados, utilizar o Docker ou Podman (`Dockerfile` / `docker-compose.yaml`). Ao rodar o comando `docker compose up` ou `podman-compose up`, tudo deverá subir, preparando o banco de dados automaticamente e a aplicação.

---

## 🛠️ Estrutura do Projeto

```plaintext
.
├── api.http
├── cmd
│   └── main.go
├── configs
│   └── config.go
├── docker
│   └── Dockerfile
├── docker-compose.yaml
├── go.mod
├── go.sum
├── init.sh
├── init.sql
├── internal
│   ├── db
│   │   ├── db.go
│   │   ├── migration
│   │   │   └── 001_create_orders_table.sql
│   │   └── repository.go
│   ├── delivery
│   │   ├── graphql
│   │   │   ├── resolver.go
│   │   │   └── schema.graphql
│   │   ├── grpc
│   │   │   ├── order.proto
│   │   │   ├── pb
│   │   │   │   ├── order_grpc.pb.go
│   │   │   │   └── order.pb.go
│   │   │   └── server.go
│   │   └── rest
│   │       └── handler.go
│   ├── domain
│   │   └── order.go
│   └── usecase
│       └── list_orders.go
├── main
├── orders.db
├── pkg
│   └── logger.go
└── README.md
```

---

## 🖥️ Tecnologias Utilizadas

- Go (Golang)
- Docker & Podman
- PostgreSQL
- gRPC
- GraphQL
- REST API

---

## 🚀 Como Executar o Projeto

### Pré-requisitos

- 🐋 Docker ou Podman instalado
- 🐳 Docker Compose ou Podman Compose instalado
- **grpcurl** (Ferramenta para interagir com o serviço gRPC)

  Instale o **grpcurl** no seu sistema (caso ainda não tenha feito isso):

  **No RHEL 9 - usado em meu computador pessoal para esse desenvolvimento**:
  ```bash
  curl -L -o grpcurl_1.9.1_linux_amd64.rpm https://github.com/fullstorydev/grpcurl/releases/download/v1.9.1/grpcurl_1.9.1_linux_amd64.rpm
  sudo rpm -ivh grpcurl_1.9.1_linux_amd64.rpm
  grpcurl --version
  ```

  **No Windows, MacOS e outras distros Linux**:
  
  [Download do grpcurl](https://github.com/fullstorydev/grpcurl) e siga as instruções para instalar.

### Passo a passo

1. Clone este repositório:
   ```bash
   git clone https://github.com/paulnune/goexpert-list-orders.git
   cd goexpert-list-orders
   ```

2. Configure o projeto Go:
   ```bash
   go mod init goexpert-list-orders
   go mod tidy
   ```

3. Inicie os containers:
   ```bash
   docker-compose up --build
   # Ou, se estiver usando Podman:
   podman-compose up --build
   ```

4. Teste as rotas disponíveis:

   - **REST API**: [http://localhost:8080/order](http://localhost:8080/order)
   - **gRPC**: Porta `50051`
   - **GraphQL**: Utilize o arquivo `api.http` para simular requests.

---

## 🧪 Testando o Projeto

1. **Teste as rotas REST**:
   Abra o arquivo `api.http` em um cliente HTTP (ex.: VSCode com a extensão REST Client) e execute as requisições para criar e listar orders.

   **Comando para criar um pedido (POST /order)**:
   ```bash
   curl -X POST http://localhost:8080/order -H "Content-Type: application/json" -d '{"customer": "Jane Doe", "total": 200.50}'
   ```

   **Comando para listar pedidos (GET /orders)**:
   ```bash
   curl -X GET http://localhost:8080/orders
   ```

   **Resposta esperada (GET /orders)**:
   ```json
   [
     {"ID": 1, "Customer": "Jane Doe", "Total": 200.5},
     {"ID": 2, "Customer": "Jane Doe", "Total": 200.5},
     {"ID": 3, "Customer": "Jane Doe", "Total": 200.5},
     {"ID": 4, "Customer": "Jane Doe", "Total": 200.5}
   ]
   ```

2. **Teste o gRPC**:

   Após compilar o `.proto`, use o **grpcurl** para fazer chamadas gRPC:
   
   - **Listar orders via gRPC**:
     ```bash
     grpcurl -plaintext -protoset order.pb -d '{}' localhost:50051 pb.OrderService/ListOrders
     ```

   **Resposta esperada (gRPC)**:
   ```json
   {
     "orders": [
       {"id": "1", "customer": "Jane Doe", "total": 200.5},
       {"id": "2", "customer": "Jane Doe", "total": 200.5},
       {"id": "3", "customer": "Jane Doe", "total": 200.5},
       {"id": "4", "customer": "Jane Doe", "total": 200.5}
     ]
   }
   ```

---

## 📝 Exemplos de Saída

### Caso Bem-sucedido

#### **POST /order**

```json
{
  "ID": 4,
  "Customer": "Jane Doe",
  "Total": 200.5
}
```

#### **GET /orders**

```json
[
  {"ID": 1, "Customer": "Jane Doe", "Total": 200.5},
  {"ID": 2, "Customer": "Jane Doe", "Total": 200.5},
  {"ID": 3, "Customer": "Jane Doe", "Total": 200.5},
  {"ID": 4, "Customer": "Jane Doe", "Total": 200.5}
]
```

---

## 👨‍💻 Autor

**Paulo Henrique Nunes Vanderley**  
- 🌐 [Site Pessoal](https://www.paulonunes.dev/)  
- 🌐 [GitHub](https://github.com/paulnune)  
- ✉️ Email: [paulo.nunes@live.de](mailto:paulo.nunes@live.de)  
- 🚀 Aluno da Pós **GoExpert 2024** pela [FullCycle](https://fullcycle.com.br)

---

## 🎉 Agradecimentos

Este repositório foi desenvolvido com muita dedicação para a **Pós GoExpert 2024**. Agradeço à equipe da **FullCycle** por proporcionar uma experiência incrível de aprendizado! 🚀

