package logic

import (
	"GoWebCode/bluebell/dao/mysql"
	"GoWebCode/bluebell/pkg/snowflake"
)

//存放业务逻辑的代码

func SignUp() {
	//1.判断用户存不存在
	mysql.QueryUserByUsername()
	//2.生成UID
	snowflake.GenID()
	//3.密码加密
	//4.保存进数据库
	mysql.InsertUser()
}
