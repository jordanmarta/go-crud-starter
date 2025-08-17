package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	rest_err "github.com/jordanmarta/go-crud-starter/src/configuration"
	"github.com/jordanmarta/go-crud-starter/src/configuration/logger"
	"github.com/jordanmarta/go-crud-starter/src/configuration/validation"
	"github.com/jordanmarta/go-crud-starter/src/controller/model/request"
	"github.com/jordanmarta/go-crud-starter/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// UpdateUser updates user information with the specified ID.
// @Summary Update User
// @Description Updates user details based on the ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be updated"
// @Param userRequest body request.UserUpdateRequest true "User information for update"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /updateUser/{userId} [put]
func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser Controller",
		zap.String("journey", "UpdateUser"),
	)

	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "UpdateUser"))

		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Userid is not a valid hex")

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error("Error trying to call UpdateUserService", err,
			zap.String("journey", "UpdateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "UpdateUser"),
	)

	c.Status(http.StatusOK)
}
