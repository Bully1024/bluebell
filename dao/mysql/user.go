package mysql

import (
	"GoWebCode/bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"

	"go.uber.org/zap"
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
		return errors.New("用户已存在")
	}
	return
}

// CheckUserExist2 检查指定用户名的用户是否存在
func CheckUserExist2(username string) (b bool) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		zap.L().Error("Check whether the user fails to be queried", zap.Error(err))
		return false
	}
	if count > 0 {
		return true
	}
	return false
}

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

// CheckPasswordisReally判断密码是否真确
func CheckPasswordisReally(user *models.User) (b bool) {
	//对输入的密码进行加密，方便校验
	h := md5.New()
	h.Write([]byte(secret))
	user.Password = hex.EncodeToString(h.Sum([]byte(user.Password)))
	//执行sql语句，查看对应username的password是否正确
	//sqlStr := `select password where username = ?`
	if 1 == 1 {
		//返回错误
		return false
	}
	return true
}
