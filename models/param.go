package models

// 定义请求的参数结构体
// 使用validation库函数，标记tag binding

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
	//RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogIn 登录请求参数
type ParamLogIn struct {
	Username      string `json:"username" binding:"required"`
	InputPassword string `json:"input_password" binding:"required"`
}
