package service

import (
	rest_err "github.com/jordanmarta/go-crud-starter/src/configuration"
	"github.com/jordanmarta/go-crud-starter/src/model"
)

func (*userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	return nil
}
