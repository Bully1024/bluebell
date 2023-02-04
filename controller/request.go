package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "user_id"

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUserID  获取当前登录用户ID
func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func GetPageInfo(c *gin.Context) (int64, int64) {
	//获取分页参数
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	var (
		size int64
		page int64
		err  error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
