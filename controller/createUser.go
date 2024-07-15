package controller

import (
	"net/http"

	"github.com/Rafaeldias0934/primeiro-crud-go.git/configuration/logger"
	"github.com/Rafaeldias0934/primeiro-crud-go.git/configuration/validation"
	"github.com/Rafaeldias0934/primeiro-crud-go.git/controller/model/request"
	"github.com/Rafaeldias0934/primeiro-crud-go.git/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}

//	err := rest_err.NewBadRequestError("Você chamou a rota errada")
//	c.JSON(err.Code, err)
