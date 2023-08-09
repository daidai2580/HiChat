package dao

import (
	"HiChat/global"
	"HiChat/models"
	"errors"
	"go.uber.org/zap"
)

func FirendList(userId uint) (*[]models.UserBasic, error) {
	relations := make([]models.Relation, 0)
	tx := global.DB.Where(" owner_id =? and type =1 ", userId).Find(&relations)
	if tx.RowsAffected == 0 {
		zap.S().Info("未查询到Relation数据")
		return nil, errors.New("未查到好友关系")
	}
	userID := make([]uint, 0)

	for _, v := range relations {
		userID = append(userID, v.TargetID)
	}

	users := make([]models.UserBasic, 0)
	tx = global.DB.Where("id in ?", userID).Find(&users)
	if tx.RowsAffected == 0 {
		zap.S().Info("未查询到Relation好友关系")
		return nil, errors.New("未查到好友")
	}
	return &users, nil
}
