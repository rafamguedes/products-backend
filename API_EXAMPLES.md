# Exemplos de Teste da API

## 1. Health Check
```bash
curl http://localhost:8080/health
```

## 2. Listar todos os produtos
```bash
curl http://localhost:8080/api/v1/products
```

## 3. Buscar produto por ID
```bash
curl http://localhost:8080/api/v1/products/1
```

## 4. Criar um novo produto
```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Produto Teste",
    "description": "Descrição do produto teste",
    "price": 99.99,
    "category": "Teste",
    "stock_quantity": 5
  }'
```

## 5. Atualizar um produto (ID 1)
```bash
curl -X PUT http://localhost:8080/api/v1/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Produto Atualizado",
    "price": 199.99,
    "stock_quantity": 15
  }'
```

## 6. Buscar produtos por categoria
```bash
curl http://localhost:8080/api/v1/products/category/Eletrônicos
```

## 7. Remover um produto (ID 10)
```bash
curl -X DELETE http://localhost:8080/api/v1/products/10
```

## Usando PowerShell (Windows)

### Health Check
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/health" -Method Get
```

### Listar produtos
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/products" -Method Get
```

### Criar produto
```powershell
$body = @{
    name = "Produto PowerShell"
    description = "Criado via PowerShell"
    price = 149.99
    category = "Teste"
    stock_quantity = 8
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/api/v1/products" -Method Post -Body $body -ContentType "application/json"
```

### Atualizar produto
```powershell
$body = @{
    name = "Produto Atualizado PS"
    price = 249.99
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/api/v1/products/1" -Method Put -Body $body -ContentType "application/json"
```

## Exemplos de Resposta

### Sucesso - Listar produtos
```json
{
  "data": [
    {
      "id": 1,
      "name": "Smartphone Samsung Galaxy S23",
      "description": "Smartphone Android com 256GB de armazenamento",
      "price": 2999.99,
      "category": "Eletrônicos",
      "stock_quantity": 50,
      "created_at": "2025-07-15T14:46:37Z",
      "updated_at": "2025-07-15T14:46:37Z"
    }
  ],
  "total": 10
}
```

### Erro - Produto não encontrado
```json
{
  "error": "Produto não encontrado"
}
```

### Erro - Dados inválidos
```json
{
  "error": "Dados inválidos",
  "details": "Key: 'CreateProductRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
}
```
