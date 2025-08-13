package main

import (
	"github.com/jordanmarta/go-crud-starter/src/controller"
	"github.com/jordanmarta/go-crud-starter/src/model/repository"
	"github.com/jordanmarta/go-crud-starter/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)

}
