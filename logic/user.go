package logic

import (
	"GoWebCode/bluebell/dao/mysql"
	"GoWebCode/bluebell/models"
	"GoWebCode/bluebell/pkg/snowflake"
)

//存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error) {
	//1.判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	//2.生成UID
	userID := snowflake.GenID()
	//构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//3.密码加密，保存进数据库
	return mysql.InsertUser(user)
}

//自己独立完成得登录逻辑处理函数
//func LogIn(p *models.ParamLogIn) (err error) {
//	if mysql.CheckUserExist2(p.Username) {
//		//用户存在，校验密码是否正确
//		//创建一个user实例，便于传参
//		user := &models.User{
//			Username: p.Username,
//			Password: p.InputPassword,
//		}
//		//校验密码是否正确
//		if mysql.CheckPasswordisReally(user) {
//			//密码正确，登录成功，返回nil
//			return
//		} else {
//			//密码错误，重新登录
//			return
//		}
//	} else {
//		zap.L().Error("Login failed：not find this user")
//		return
//	}
//}

func LogIn(p *models.ParamLogIn) (err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.InputPassword,
	}
	return mysql.Login(user)
}
