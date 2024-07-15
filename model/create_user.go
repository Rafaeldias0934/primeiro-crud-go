package model

import (
	"fmt"

	"github.com/Rafaeldias0934/primeiro-crud-go.git/configuration/logger"
	"github.com/Rafaeldias0934/primeiro-crud-go.git/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("Journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil

}
