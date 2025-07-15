# API REST de Produtos em Go

Uma API REST completa para gerenciamento de produtos, construída com Go, Gin Framework e PostgreSQL.

## 🚀 Funcionalidades

- ✅ CRUD completo de produtos
- ✅ Busca por categoria
- ✅ Validação de dados
- ✅ Tratamento de erros
- ✅ Banco de dados PostgreSQL
- ✅ Docker Compose para ambiente de desenvolvimento
- ✅ Dados de exemplo (seed data)

## 📋 Pré-requisitos

- Go 1.21 ou superior
- Docker e Docker Compose
- Git

## 🛠️ Instalação e Execução

### 1. Clone o repositório
```bash
git clone <seu-repositorio>
cd apiRestGolang
```

### 2. Inicie o banco de dados PostgreSQL
```bash
docker-compose up -d
```

### 3. Instale as dependências
```bash
go mod tidy
```

### 4. Execute a aplicação
```bash
go run main.go
```

A API estará disponível em: `http://localhost:8080`

## 📚 Endpoints da API

### Health Check
- `GET /health` - Verifica o status da API

### Produtos
- `GET /api/v1/products` - Lista todos os produtos
- `GET /api/v1/products/:id` - Busca produto por ID
- `POST /api/v1/products` - Cria um novo produto
- `PUT /api/v1/products/:id` - Atualiza um produto
- `DELETE /api/v1/products/:id` - Remove um produto
- `GET /api/v1/products/category/:category` - Lista produtos por categoria

## 📄 Exemplos de Uso

### Criar um produto
```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Novo Produto",
    "description": "Descrição do produto",
    "price": 199.99,
    "category": "Eletrônicos",
    "stock_quantity": 10
  }'
```

### Listar todos os produtos
```bash
curl http://localhost:8080/api/v1/products
```

### Buscar produto por ID
```bash
curl http://localhost:8080/api/v1/products/1
```

### Atualizar um produto
```bash
curl -X PUT http://localhost:8080/api/v1/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Produto Atualizado",
    "price": 299.99
  }'
```

### Remover um produto
```bash
curl -X DELETE http://localhost:8080/api/v1/products/1
```

### Buscar produtos por categoria
```bash
curl http://localhost:8080/api/v1/products/category/Eletrônicos
```

## 🗄️ Estrutura do Banco de Dados

### Tabela: products
```sql
id              SERIAL PRIMARY KEY
name            VARCHAR(255) NOT NULL
description     TEXT
price           DECIMAL(10,2) NOT NULL
category        VARCHAR(100)
stock_quantity  INTEGER DEFAULT 0
created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
```

## 📁 Estrutura do Projeto

```
apiRestGolang/
├── main.go                           # Arquivo principal
├── go.mod                           # Dependências do Go
├── go.sum                           # Checksums das dependências
├── docker-compose.yml               # Configuração do Docker
├── init.sql                         # Script de inicialização do DB
├── .env                             # Variáveis de ambiente
├── database/
│   └── connection.go                # Conexão com o banco
├── models/
│   └── product.go                   # Modelos de dados
├── repositories/
│   └── product_repository.go       # Operações de banco de dados
├── handlers/
│   └── product_handler.go          # Controladores da API
└── routes/
    └── routes.go                    # Configuração das rotas
```

## 🧪 Dados de Teste

O banco é inicializado com 10 produtos de exemplo em diferentes categorias:
- Eletrônicos
- Computadores
- Áudio
- Acessórios
- Monitores
- Fotografia
- Armazenamento

## 🐳 Docker

Para parar o banco de dados:
```bash
docker-compose down
```

Para remover os volumes (dados serão perdidos):
```bash
docker-compose down -v
```

## 🔧 Configuração

Edite o arquivo `.env` para personalizar as configurações:
- Porta do servidor
- Configurações do banco de dados
- Modo do Gin (debug/release)

## 📝 Notas

- A API inclui middleware de CORS para desenvolvimento
- Os logs são habilitados por padrão
- A validação de dados é feita automaticamente pelo Gin
- Todos os endpoints retornam JSON
- Tratamento de erros adequado em todas as operações
