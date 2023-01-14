package controller

import (
	"GoWebCode/bluebell/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//社区模块

func CommunityHandler(c *gin.Context) {
	//查询到所有的社区(community_id,community_name)以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不要轻易把服务端报错返回给外面
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	//1.获取社区id
	idStr := c.Param("id")
	//从URL获取参数 类型为字符串，需要做类型转换
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2.根据ID查询社区详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不要轻易把服务端报错返回给外面
		return
	}
	ResponseSuccess(c, data)
}
