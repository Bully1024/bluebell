package models

// 定义请求的参数结构体
// 使用validation库函数，标记tag

type ParamSignUp struct {
	Username   string `form:"username" json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}
