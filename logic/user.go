package logic

import (
	"GoWebCode/bluebell/dao/mysql"
	"GoWebCode/bluebell/models"
	"GoWebCode/bluebell/pkg/snowflake"
	"errors"
)

//存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) {
	//1.判断用户存不存在
	if mysql.CheckUserExist(p.Username) {
		return errors.New("用户已存在")
	}
	//2.生成UID
	userID := snowflake.GenID()
	//构造一个User实例
	//3.密码加密
	//4.保存进数据库
	mysql.InsertUser()
}
