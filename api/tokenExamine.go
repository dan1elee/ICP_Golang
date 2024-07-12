package api

import (
	"ICP_Golang/token"
	"errors"

	"github.com/gin-gonic/gin"
)

func tokenValidation(c *gin.Context, level uint) error {
	userToken, exist := c.GetPostForm("token")
	if !exist {
		return errors.New("token field not exist")
	}
	tokenStruct, err := token.ParseToken(userToken)
	if err != nil {
		return err
	}
	if !token.HaveAccess(tokenStruct, level) {
		return errors.New("access wrong")
	}
	return nil
}
