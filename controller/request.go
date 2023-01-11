package controller

import (
	"GoWebCode/bluebell/middlewares"
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrorUserBotLogin = errors.New("用户未登录")

// GetCurrentUser 获取当前登录用户ID
func GetCurrentUser(c gin.Context) (userID int64, err error) {
	uid, ok := c.Get(middlewares.CtxUserIDKey)
	if !ok {
		err = ErrorUserBotLogin
		return
	}

	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserBotLogin
		return
	}
	return

}
