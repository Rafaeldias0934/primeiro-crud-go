package routes

import (
	"github.com/Rafaeldias0934/primeiro-crud-go.git/controller"
	"github.com/gin-gonic/gin"
)

// função criada para inicializar as rotas
func InirRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById:userId", controller.FindUserById)
	r.GET("/getUserByEmail:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser:userId", controller.UpdateUser)
	r.DELETE("/deleteUser:userId", controller.DeleteUser)

}
