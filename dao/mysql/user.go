package mysql

// 把每一步数据库操作封装成函数，等待logic层根据业务需求调用

func CheckUserExist(username string) bool {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	db.Get(&count, sqlStr, username)
	return
}
func InsertUser() {
	//
	//执行SQL语句入库
}
