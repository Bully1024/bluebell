package logic

import (
	"GoWebCode/bluebell/dao/mysql"
	"GoWebCode/bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	//查询数据 查找到所有的community并返回
	return mysql.GetCommunityList()
}

// GetCommunityDetail 分类详情
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
