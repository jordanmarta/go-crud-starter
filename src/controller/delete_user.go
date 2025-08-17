package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	rest_err "github.com/jordanmarta/go-crud-starter/src/configuration"
	"github.com/jordanmarta/go-crud-starter/src/configuration/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// DeleteUser deletes a user with the specified ID.
// @Summary Delete User
// @Description Deletes a user based on the ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be deleted"
// @Success 200
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /deleteUser/{userId} [delete]
func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser Controller",
		zap.String("journey", "DeleteUser"),
	)

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Userid is not a valid hex")

		c.JSON(errRest.Code, errRest)
		return
	}

	err := uc.service.DeleteUser(userId)

	if err != nil {
		logger.Error("Error trying to call DeleteUserService", err,
			zap.String("journey", "DeleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "DeleteUser"),
	)

	c.Status(http.StatusOK)
}
