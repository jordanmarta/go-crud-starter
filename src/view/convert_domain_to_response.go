package view

import (
	"github.com/jordanmarta/go-crud-starter/src/controller/model/response"
	"github.com/jordanmarta/go-crud-starter/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
