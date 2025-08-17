package main

import (
	"context"
	"log"

	_ "github.com/jordanmarta/go-crud-starter/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	mongodb "github.com/jordanmarta/go-crud-starter/src/configuration/database"
	"github.com/jordanmarta/go-crud-starter/src/configuration/logger"
	"github.com/jordanmarta/go-crud-starter/src/controller/routes"
)

// @title Meu Primeiro CRUD em Go | HunCoding
// @version 1.0
// @description API for crud operations on users
// @host localhost:8080
// @BasePath /
// @schemes http
// @license MIT
func main() {

	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
