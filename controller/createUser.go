package controller

import (
	"net/http"

	"github.com/Rafaeldias0934/primeiro-crud-go.git/configuration/logger"
	"github.com/Rafaeldias0934/primeiro-crud-go.git/configuration/validation"
	"github.com/Rafaeldias0934/primeiro-crud-go.git/controller/model/request"
	"github.com/Rafaeldias0934/primeiro-crud-go.git/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {

	logger.Info("Init CreatUser controller",
		zap.String("journey", "CreateUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		//c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)
}

//	err := rest_err.NewBadRequestError("Você chamou a rota errada")
//	c.JSON(err.Code, err)
