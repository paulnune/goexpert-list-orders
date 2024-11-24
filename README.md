
# 🚀 Desafio GoExpert - Listagem de Orders

Este projeto foi desenvolvido por **Paulo Nunes**.

---

## 📝 Sobre o Desafio

Olá devs! Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders. Esta listagem precisa ser feita com:

- 📡 **Endpoint REST** (GET `/order`)
- 🔗 **Service ListOrders com GRPC**
- 🧩 **Query ListOrders GraphQL**

Não esqueça de criar as migrações necessárias e o arquivo `api.http` com as requests para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker ou Podman (`Dockerfile` / `docker-compose.yaml`). Ao rodar o comando `docker compose up` ou `podman-compose up`, tudo deverá subir, preparando o banco de dados automaticamente.

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

   - REST API: [http://localhost:8080/order](http://localhost:8080/order)
   - gRPC: Porta `50051`
   - GraphQL: Utilize o arquivo `api.http` para simular requests.

---

## 🧪 Testando o Projeto

1. Abra o arquivo `api.http` em um cliente HTTP (ex.: VSCode com a extensão REST Client).
2. Execute as requisições para criar e listar orders.

---

## 📝 Exemplos de Saída

### Caso Bem-sucedido

```
Pedidos:
ID: 1, Cliente: John Doe, Total: 123.45
Fim da listagem de pedidos.
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
