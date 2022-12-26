package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luankkobs/go-crud/src/configuration/logger"
	"github.com/luankkobs/go-crud/src/configuration/validation"
	"github.com/luankkobs/go-crud/src/controller/model/request"
	"github.com/luankkobs/go-crud/src/controller/model/response"
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

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully",
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		})

	c.JSON(http.StatusOK, response)
}
