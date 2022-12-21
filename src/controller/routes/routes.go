package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luankkobs/go-crud/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUserById/:userId", controller.DeleteUser)
}
