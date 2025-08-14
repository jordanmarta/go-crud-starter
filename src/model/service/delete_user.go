package service

import (
	rest_err "github.com/jordanmarta/go-crud-starter/src/configuration"
	"github.com/jordanmarta/go-crud-starter/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Init deleteUser model", zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUser(userId)

	if err != nil {
		return err
	}

	return nil
}
