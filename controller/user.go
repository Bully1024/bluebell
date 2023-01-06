package controller

import (
	"GoWebCode/bluebell/logic"
	_ "GoWebCode/bluebell/logic"
	"GoWebCode/bluebell/models"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误，直接返回响应 只能判断参数的类型!!!
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msgs": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msgg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})
		return
	}
	//手动对请求参数进行详细的业务规则校验 使用gin框架内置库validator做参数校验，就不需要手动进行校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 {
	//	zap.L().Error("SignUp with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}
	fmt.Println(p)
	//2.业务处理
	logic.SignUp(p)
	//3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
