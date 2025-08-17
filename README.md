# go-crud-starter

A minimal Go REST API with CRUD operations for users, backed by MongoDB and Docker-ready.

---

## Technologies & Libraries

- Go 1.18+  
- gin-gonic/gin              (HTTP router)  
- go.mongodb.org/mongo-driver (MongoDB client)  
- go.uber.org/zap            (structured logging)  
- github.com/go-playground/validator/v10 (request validation)  
- github.com/swaggo/http-swagger (Swagger UI)  
- github.com/stretchr/testify  (unit testing)

---

## Prerequisites

- Go
- Docker

---

## Quick Start

1. Clone the repo  
   ```bash
   git clone https://github.com/jordanmarta/go-crud-starter.git
   cd go-crud-starter
   ```

2. Launch MongoDB and the app  
   ```bash
   docker compose up
   go run main.go
   ```

3. Open in your browser  
   ```
   http://localhost:8080
   ```

---

## API Endpoints

- **POST**   /createUser  
- **GET**    /getUserById/{userId}  
- **GET**    /getUserByEmail/{userEmail}  
- **PUT**    /updateUser/{userId}  
- **DELETE** /deleteUser/{userId}  
- **POST**   /login  

Explore and test via Swagger UI:  
```
http://localhost:8080/swagger/index.html
```

---

## Credit

This project follows the HunCoding tutorial series:  
https://www.youtube.com/watch?v=vxDqv6BKZCw&list=PLm-xZWCprwYQ3gyCxJ8TR1L2ZnUOPvOpr  

Original repository by HunCoding:  
https://github.com/HunCoding/meu-primeiro-crud-go

---

