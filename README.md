# go-crud-starter

Uma API REST minimalista em Go com operações CRUD para usuários, utilizando MongoDB e pronta para Docker.

---

## Tecnologias & Bibliotecas

- Go 1.18+  
- gin-gonic/gin (roteador HTTP)  
- go.mongodb.org/mongo-driver (cliente MongoDB)  
- go.uber.org/zap (log estruturado)  
- github.com/go-playground/validator/v10 (validação de requisições)  
- github.com/swaggo/http-swagger (interface Swagger UI)  
- github.com/stretchr/testify (testes unitários)

---

## Pré-requisitos

- Go  
- Docker  

---

## Início Rápido

1. Clone o repositório  
   ```bash
   git clone https://github.com/jordanmarta/go-crud-starter.git
   cd go-crud-starter
   ```

2. Inicie o MongoDB e a aplicação  
   ```bash
   docker compose up
   go run main.go
   ```

3. Acesse no navegador  
   ```
   http://localhost:8080
   ```

---

## Endpoints da API

- **POST**   /createUser  
- **GET**    /getUserById/{userId}  
- **GET**    /getUserByEmail/{userEmail}  
- **PUT**    /updateUser/{userId}  
- **DELETE** /deleteUser/{userId}  
- **POST**   /login  

Explore e teste via Swagger UI:  
```
http://localhost:8080/swagger/index.html
```

---

## Créditos

Este projeto foi baseado na série de tutoriais do canal HunCoding:  
https://www.youtube.com/watch?v=vxDqv6BKZCw&list=PLm-xZWCprwYQ3gyCxJ8TR1L2ZnUOPvOpr  

Repositório original do HunCoding:  
https://github.com/HunCoding/meu-primeiro-crud-go

---
