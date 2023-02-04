package logic

import (
	"GoWebCode/bluebell/dao/mysql"
	"GoWebCode/bluebell/models"
	"GoWebCode/bluebell/pkg/jwt"
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

// 自己独立完成得登录逻辑处理函数 -diy
//func LogIn(p *models.ParamLogIn) (err error) {
//	//用户存在，校验密码是否正确
//	//创建一个user实例，便于传参
//	user := &models.User{
//		Username: p.Username,
//		Password: p.InputPassword,
//	}
//	//校验密码是否正确
//	return mysql.CheckPasswordisReally(user)
//}

func LogIn(p *models.ParamLogIn) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.InputPassword,
	}
	//传递的是指针，就能达到user.UserID
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	//生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
