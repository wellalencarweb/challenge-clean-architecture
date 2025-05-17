# 🧱 Challenge Clean Architecture (Go)

Este projeto demonstra a aplicação dos princípios da Clean Architecture com foco em modularidade, testabilidade e independência de frameworks, utilizando a linguagem Go.

## 🧩 Tecnologias

- Go (Golang)
- PostgreSQL (via Docker)
- REST API (padrão HTTP)
- GraphQL API
- gRPC API
- Docker & Docker Compose
- Testes automatizados (unitários)
- Evans (gRPC CLI)
- VSCode REST Client (`api.http`)

## 🚀 Como executar

### 1. Subir os containers (MySQL e RabbitMQ)

```bash
make up
```
resultado esperado no terminal:
```bash
[+] Running 3/3
 ✔ Network challenge-clean-architecture_default  Created
 ✔ Container mysql                               Started
 ✔ Container rabbitmq                            Started
```

### 2. Instalar as dependências do projeto

```bash
make build
```

### 3. Executar o sistema de pedidos

```bash
make run
```

Ao executar o sistema, você verá a seguinte saída:

```bash
Starting web server on port :8000
Starting gRPC server on port 50051
Starting GraphQL server on port 8080
```

### Observações

- O servidor HTTP está disponível na porta **8000**
- O servidor gRPC está disponível na porta **50051**
- O servidor GraphQL está disponível na porta **8080**

---

## 📁 Estrutura de pastas

```
.
├── api                 # Arquivos REST client (`api.http`)
├── cmd                 # Entry point da aplicação
├── configs             # Configurações (ex: .env, docker, etc)
├── internal            # Domínio e regras de negócio
│   ├── domain          # Entidades e interfaces
│   ├── application     # Casos de uso
│   ├── infrastructure  # Repositórios, banco, etc
├── pkg                 # Pacotes compartilhados/utilitários
├── sql                 # Scripts SQL de migração
```

---

## 🔌 Integrações disponíveis

### ✅ REST (via HTTP Client)

Utilize o arquivo `api/api.http` no VSCode com a extensão **REST Client** para testar os endpoints da API REST.

#### 📤 Criar um pedido

```http
POST http://localhost:8000/order HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
  "id": "TOM",
  "price": 13.5,
  "tax": 0.8
}
```

#### 📥 Listar pedidos

```http
GET http://localhost:8000/order HTTP/1.1
Content-Type: application/json
```

A resposta será:

```json
{
  "Orders": [
    { "id": "ABC", "price": 22.4, "tax": 0.6, "final_price": 23 },
    { "id": "DEF", "price": 23.4, "tax": 0.6, "final_price": 24 },
    { "id": "TOM", "price": 10.5, "tax": 0.5, "final_price": 11 }
  ]
}
```
![api.png](/img/api.png)

---

### ✅ GraphQL (porta 8080)

Acesse via navegador o [GraphQL Playground](http://localhost:8080).

#### Query de exemplo:

```graphql
query queryOrders {
  orders {
    id
    FinalPrice
    Price
    Tax
  }
}
```

A resposta será:

```json
{
  "data": {
    "orders": [
      { "id": "ABC", "FinalPrice": 23, "Price": 22.4, "Tax": 0.6 },
      { "id": "DEF", "FinalPrice": 24, "Price": 23.4, "Tax": 0.6 },
      { "id": "TOM", "FinalPrice": 11, "Price": 10.5, "Tax": 0.5 }
    ]
  }
}
```
![graphql.png](/img/graphql.png)

---

### ✅ gRPC (porta 50051)

Use o [Evans CLI](https://github.com/ktr0731/evans) para interagir via terminal.

```bash
evans -r repl
```

Comandos:

```evans
package pb
service OrderService
call ListOrders
```

Resposta:

```json
{
  "orders": [
    { "id": "ABC", "price": 22.4, "tax": 0.6, "finalPrice": 23 },
    { "id": "DEF", "price": 23.4, "tax": 0.6, "finalPrice": 24 },
    { "id": "GHI", "price": 10.5, "tax": 0.5, "finalPrice": 11 }
  ]
}
```
![evans_rpc.png](/img/evans_grpc.png)

---

## 🧪 Testes

Todos os casos de uso estão cobertos por testes unitários. Execute com:

```bash
make test
```

---


## 🛠️ Comandos disponíveis no Makefile

| Comando     | Descrição                                      |
|-------------|------------------------------------------------|
| `make up`   | Sobe os containers Docker (MySQL, RabbitMQ)    |
| `make down` | Para e remove os containers Docker             |
| `make build`| Organiza e baixa as dependências do projeto    |
| `make run`  | Executa a aplicação (inicia os servidores)     |
| `make test` | Executa todos os testes unitários com verbose  |

---

## 📦 Build

```bash
go build -o bin/app cmd/main.go
```

---

## 🧼 Padrão arquitetural

Este projeto segue os princípios da **Clean Architecture**:
- Independência de frameworks
- Testabilidade
- Separação clara de camadas
- Inversão de dependências

---

## 📄 Licença

MIT License