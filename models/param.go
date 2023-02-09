package models

// 定义请求的参数结构体
// 使用validation库函数，标记tag binding

const (
	DefaultPage = 1
	DefaultSize = 10
	Ordertime   = "time"
	OrderScore  = "score"
)

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

// ParamVoteData 投票数据
type ParamVoteData struct {
	//UserID 直接从当前用户请求中获取
	PostID    string `json:"post_id" binding:"required"`       //帖子id
	Direction int8   `json:"direction" binding:"oneof=1 0 -1"` //赞成票：1，反对票：-1 取消投票：0
}

// ParamPostList 获取帖子列表参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"` //可以为空
	Page        int64  `json:"page" form:"page"`
	Size        int64  `json:"size" form:"size"`
	Order       string `json:"order" form:"order"`
}
