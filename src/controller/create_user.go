package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luankkobs/go-crud/src/configuration/logger"
	"github.com/luankkobs/go-crud/src/configuration/validation"
	"github.com/luankkobs/go-crud/src/controller/model/request"
	"github.com/luankkobs/go-crud/src/model"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		})
	c.String(http.StatusOK, "")
}
