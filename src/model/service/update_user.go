package service

import (
	rest_err "github.com/jordanmarta/go-crud-starter/src/configuration"
	"github.com/jordanmarta/go-crud-starter/src/configuration/logger"
	"github.com/jordanmarta/go-crud-starter/src/model"
	"go.uber.org/zap"
)

func (ud userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)

	if err != nil {
		return err
	}

	return nil
}
