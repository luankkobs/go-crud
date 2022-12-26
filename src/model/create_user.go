package model

import (
	"fmt"
	"github.com/luankkobs/go-crud/src/configuration/logger"
	"github.com/luankkobs/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()
	fmt.Println(ud)
	return nil
}
