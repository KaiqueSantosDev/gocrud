# Product API

## Como Usar:
    1 - Instale go: go.dev
    2 - Crie arquivo .env como no exemplo
    3 - Execute a função main.go usando go run main.go


## Endpoints:


### Pegar Todos os Produtos:
```
curl "http://localhost:8000/" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```


### Criar Produto:
```
curl -X "POST" "http://localhost:8000/" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{   
        "name": "Teste",
        "description":"teste",
        "price": 28
        }'
```

### Pegar Produto pelo ID:
```
curl http://localhost:8000/1
```

### Atualizar Produto:
```
curl -X "PUT" "http://localhost:8000/1" \
     -H 'Accept: application/json' \
     -H 'Content-Type: application/json' \
     -d $'{
  "name": "Produto1",
  "description": "Atualizado",
  "price": 15
}'

```

### Deletar Produto:
```
curl -X "DELETE" "http://localhost:8000/2" \
     -H 'Accept: application/json' \
     -H 'Content-Type: application/json'
```


### Compare Produto:
```
curl -XPOST -H "Content-type: application/json" -d '[
    {   
        "name": "Teste",
        "description":"teste",
        "price": 28
        },
    {
        "name": "Teste",
        "description":"teste",
        "price": 25
    }

]' 'localhost:8000/compare/'
```