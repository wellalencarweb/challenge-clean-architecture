# **Desafio de Clean Architecture — Pós Go Expert (Full Cycle)**

Este repositório contém a solução para o desafio proposto na disciplina de Clean Architecture da **Pós-graduação Go Expert**, promovida pela **Full Cycle**.

## ✅ Objetivo

O desafio consiste em implementar uma funcionalidade de listagem de pedidos (**orders**) com base nos princípios da **Clean Architecture**, utilizando múltiplas interfaces de comunicação:

- **REST API**: `GET /order`
- **gRPC Service**: `ListOrders`
- **GraphQL Query**: `ListOrders`

Além disso, é necessário utilizar Docker para provisionamento do ambiente de banco de dados.

---

## 🚀 Tecnologias e Estratégias Utilizadas

- **Clean Architecture**: Separação de responsabilidades entre camadas (usecase, entity, interface).
- **gRPC + Protocol Buffers**: Comunicação eficiente entre serviços.
- **GraphQL (gqlgen)**: Query estruturada e flexível.
- **REST API (chi)**: Interface simples e comum para integração.
- **Wire (Google)**: Injeção de dependências automatizada.
- **Docker / Docker Compose**: Gerenciamento de infraestrutura (MySQL, RabbitMQ).
- **SQL de inicialização**: Criação automática de tabelas via `initdb.sql`.

---

## 📦 Inicialização do Projeto

Após clonar o repositório, siga os passos abaixo:

### 1. Suba os containers com Docker Compose:

```bash
make up
```

Isso irá levantar o MySQL (com a tabela `orders` já criada) e o RabbitMQ.

### 2. Rode a aplicação:

```bash
make tidy
make run
```

A aplicação estará disponível nos seguintes serviços:

- 🌐 Web Server (REST): `http://localhost:8000`
- ⚙️ gRPC Server: `localhost:50051`
- 🔍 GraphQL Playground: `http://localhost:8080`

---

## 🧪 Chamadas e Testes

### REST API

Use o arquivo `api/api.http` (no VSCode com REST Client) para testar as rotas:

```
GET http://localhost:8000/order
```

### gRPC

Utilize o **Evans** para chamadas interativas no terminal:

```bash
evans --host localhost --port 50051 -r repl
```

### GraphQL

Acesse via navegador:

```
http://localhost:8080
```

Exemplo de query:

```graphql
query {
  orders {
    id
    FinalPrice
    Price
    Tax
  }
}
```

---

## 🛠️ Comandos via Makefile

| Comando         | Descrição                                       |
|-----------------|-------------------------------------------------|
| `make up`       | Sobe os containers Docker (MySQL, RabbitMQ)     |
| `make run`      | Roda a aplicação principal                      |
| `make build`    | Compila o binário em `./bin/ordersystem`        |
| `make test`     | Executa os testes unitários                     |
| `make proto`    | Compila os arquivos `.proto`                    |
| `make tidy`     | Atualiza dependências com `go mod tidy`         |
| `make clean`    | Remove arquivos gerados (`bin/`)                |

---

## 📁 Estrutura do Projeto

```
challenge-clean-architecture/
├── api/                    # Interfaces: REST, GraphQL, HTTP
├── cmd/ordersystem/       # Entrypoint da aplicação
├── internal/
│   ├── entity/             # Entidades de domínio
│   ├── usecase/            # Casos de uso (ex: ListOrders)
│   └── infra/              # Infraestrutura: DB, MQ, etc
├── proto/                 # Arquivos gRPC (.proto)
├── script/                # Scripts (ex: initdb.sql)
├── .img/                  # Imagens para o README
├── docker-compose.yml
├── Makefile
├── go.mod
└── README.md
```

---

## 📚 Referências

- [go-chi/chi](https://github.com/go-chi/chi)
- [spf13/viper](https://github.com/spf13/viper)
- [gqlgen](https://gqlgen.com/)
- [gRPC](https://grpc.io)
- [Protocol Buffers](https://protobuf.dev/)
- [Evans (gRPC CLI)](https://github.com/ktr0731/evans)
- [Google Wire](https://github.com/google/wire)
