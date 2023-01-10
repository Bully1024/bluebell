package controller

import (
	"GoWebCode/bluebell/dao/mysql"
	"GoWebCode/bluebell/logic"
	_ "GoWebCode/bluebell/logic"
	"GoWebCode/bluebell/models"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(models.ParamSignUp)
	//var p models.ParamSignUp
	if err := c.ShouldBind(&p); err != nil {
		//请求参数有误，直接返回响应 只能判断参数的类型!!!
		fmt.Println(err)
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": err.Error(),
			//})
			//优化：使用封装函数
			ResponseError(c, CodeInvalidParam)
			return
		}
		//Todo 无法翻译中文 login函数存在同样问题
		// validator.ValidationErrors类型错误则进行翻译
		//c.JSON(http.StatusOK, gin.H{
		//	"msgg": removeTopStruct(errs.Translate(trans)), //翻译错误
		//})
		//优化：使用封装函数
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
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
	//fmt.Println(p)
	//2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.Signup failed", zap.Error(err))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "注册失败",
		//})
		//优化，使用封装函数
		//Todo errors.Is 1.13新特性
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "success",
	//})
	//优化：封装函数
	ResponseSuccess(c, nil)
}

// LoginHandler 登录函数
//func LoginHandler(c *gin.Context) {
//	//1.获得请求参数
//	p := new(models.ParamLogIn)
//	if err := c.ShouldBind(p); err != nil {
//		fmt.Println(err)
//		zap.L().Error("LogIn with invalid param", zap.Error(err))
//		c.JSON(http.StatusOK, gin.H{
//			"msg": err.Error(),
//		})
//		return
//	}
//	//2.校验参数
//	if err := logic.LogIn(p); err != nil {
//		zap.L().Error("logic.Login failed", zap.Error(err))
//		c.JSON(http.StatusOK, gin.H{
//			"msg": "登录失败",
//		})
//		return
//	}
//	//3.业务逻辑处理
//	//4.返回响应
//	c.JSON(http.StatusOK, gin.H{
//		"msg": "success",
//	})
//}

/*
自己写完整的登录过程的问题
zap日志使用不清楚
sql语句不会判断密码是否相等
*/

func LoginHandler(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(models.ParamLogIn)
	if err := c.ShouldBind(p); err != nil {
		//请求参数有误，直接返回响应 只能判断参数的类型!!!
		zap.L().Error("Login with invalid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		// validator.ValidationErrors类型错误则进行翻译
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		//})
		return
	}
	//2.业务逻辑处理
	if err := logic.LogIn(p); err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "用户名或密码错误",
		//})
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//3.返回响应
	//c.JSON(http.StatusOK, gin.H{
	//	"mag": "登录成功",
	//})
	ResponseSuccess(c, nil)
}
