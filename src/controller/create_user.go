package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luankkobs/go-crud/src/configuration/validation"
	"github.com/luankkobs/go-crud/src/controller/model/request"
	"log"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to bind user request: %s", err.Error())
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)

}
