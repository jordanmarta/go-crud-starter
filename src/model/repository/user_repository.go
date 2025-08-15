package repository

import (
	rest_err "github.com/jordanmarta/go-crud-starter/src/configuration"
	"github.com/jordanmarta/go-crud-starter/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (
		model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr

	DeleteUser(userId string) *rest_err.RestErr

	FindUserByEmail(email string) (
		model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailAndPassword(email, password string) (
		model.UserDomainInterface, *rest_err.RestErr)

	FindUserByID(email string) (
		model.UserDomainInterface, *rest_err.RestErr)
}
