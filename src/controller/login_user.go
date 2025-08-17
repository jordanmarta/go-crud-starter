package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordanmarta/go-crud-starter/src/configuration/logger"
	"github.com/jordanmarta/go-crud-starter/src/configuration/validation"
	"github.com/jordanmarta/go-crud-starter/src/controller/model/request"
	"github.com/jordanmarta/go-crud-starter/src/model"
	"github.com/jordanmarta/go-crud-starter/src/view"
	"go.uber.org/zap"
)

// LoginUser allows a user to log in and obtain an authentication token.
// @Summary User Login
// @Description Allows a user to log in and receive an authentication token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param userLogin body request.UserLogin true "User login credentials"
// @Success 200 {object} response.UserResponse "Login successful, authentication token provided"
// @Header 200 {string} Authorization "Authentication token"
// @Failure 403 {object} rest_err.RestErr "Error: Invalid login credentials"
// @Router /login [post]
func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser Controller",
		zap.String("journey", "loginUser"),
	)

	var userLogin request.UserLogin

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "loginUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(
		userLogin.Email,
		userLogin.Password,
	)

	domainResult, token, err := uc.service.LoginUserServices(domain)

	if err != nil {
		logger.Error("Error trying to call loginUserService", err,
			zap.String("journey", "loginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "loginUser"),
	)

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
