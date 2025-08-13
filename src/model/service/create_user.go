package service

import (
	rest_err "github.com/jordanmarta/go-crud-starter/src/configuration"
	"github.com/jordanmarta/go-crud-starter/src/configuration/logger"
	"github.com/jordanmarta/go-crud-starter/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}

	return userDomainRepository, nil
}
