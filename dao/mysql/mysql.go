package mysql

import (
	"GoWebCode/bluebell/settings"
	"fmt"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Todo 视频用的是sqlx，重构为sql

// 定义一个全局对象db
var db *sqlx.DB

// Init 定义一个初始化数据库的函数
func Init(cfg *settings.MySQLConfig) (err error) {
	// DSN:Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)
	//Todo 博客
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	//db, err = sql.Open("mysql", dsn)
	//if err != nil {
	//	return err
	//}
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
	}
	//fmt.Println(viper.GetInt("mysql.max_open_conns"))
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}
