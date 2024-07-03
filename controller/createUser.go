package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rafaeldias0934/primeiro-crud-go.git/configuration/validation"
	"github.com/Rafaeldias0934/primeiro-crud-go.git/controller/model/request"
	"github.com/Rafaeldias0934/primeiro-crud-go.git/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	log.Println("Init CreatUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Erro trying to marshal object, error =%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		//c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(http.StatusOK, response)
}

//	err := rest_err.NewBadRequestError("Você chamou a rota errada")
//	c.JSON(err.Code, err)
