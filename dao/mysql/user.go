package mysql

import (
	"GoWebCode/bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

// 把每一步数据库操作封装成函数，等待logic层根据业务需求调用

const secret = "liwenzhou.com"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// CheckUserExist2 检查指定用户名的用户是否存在 -diy 合并至mysql.login 弃用
//func CheckUserExist2(username string) (b bool) {
//	sqlStr := `select count(user_id) from user where username = ?`
//	var count int
//	if err := db.Get(&count, sqlStr, username); err != nil {
//		zap.L().Error("Check whether the user fails to be queried", zap.Error(err))
//		return false
//	}
//	if count > 0 {
//		return true
//	}
//	return false
//}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	//对密码进行加密
	user.Password = encryptPassword(user.Password)
	//执行SQL语句入库
	sqlStr := `insert into user(user_id,username,password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// encryptPassword加密密码
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// CheckPasswordisReally 判断密码是否真确 -diy
//func CheckPasswordisReally(user *models.User) (err error) {
//	//对输入的密码进行加密，方便校验
//	//h := md5.New()
//	//h.Write([]byte(secret))
//	//user.Password = hex.EncodeToString(h.Sum([]byte(user.Password)))
//
//	//用户输入的密码
//	user.Password = encryptPassword(user.Password)
//	//执行sql语句，查看对应username的password是否正确
//	sqlStr := `select password from user where username = ?`
//	var rPassword string
//	err = db.Get(rPassword, sqlStr, user.Username)
//	fmt.Println("mysql real:", rPassword)
//	fmt.Println("input:", user.Password)
//	if err == sql.ErrNoRows {
//		return errors.New("用户或密码错误")
//	}
//	if err != nil {
//		//查询数据库错误
//		return err
//	}
//	if user.Password != rPassword {
//		return errors.New("密码错误")
//	}
//	return
//}

func Login(user *models.User) (err error) {
	oPassword := user.Password //用户登录Input_password
	sqlStr := `select user_id,username,password from user where username= ?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		//查询数据库错误
		return err
	}
	//判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}
