package controller

import (
	"GoWebCode/bluebell/logic"
	"GoWebCode/bluebell/models"
	"fmt"

	"go.uber.org/zap"

	"github.com/go-playground/validator"

	"github.com/gin-gonic/gin"
)

// 投票

type VoteData struct {
	//UserID 直接从当前用户请求中获取
	PostID    int64 `json:"post_id,string"`   //帖子id
	Direction int   `json:"direction,string"` //赞成票：1，反对票：-1
}

func PostVoteController(c *gin.Context) {
	//参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		//fmt.Println("bool", err.(validator.ValidationErrors))
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			fmt.Println("trans errs after:", errs)
			ResponseError(c, CodeInvalidParam)
			return
		}

		errData := removeTopStruct(errs.Translate(trans)) //removeTopStruct去除提示信息中结构体的名称,翻译
		fmt.Println("trans after errData is:", errData)
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	//获取当前用户ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	//具体投票业务逻辑
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
