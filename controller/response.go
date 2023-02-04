package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	"code":10001,	//程序中的错误码
	“msg":xx,		//提示信息
	"data":{}		//数据
*/

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code ResCode) {
	//也可以使用g.H，其本质就是一个map
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Mag(),
		Data: nil,
	})
}
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	//也可以使用g.H，其本质就是一个map
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Mag(),
		Data: nil,
	})
}
func ResponseSuccess(c *gin.Context, data interface{}) {
	//也可以使用g.H，其本质就是一个map
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Mag(),
		Data: data,
	})
}
