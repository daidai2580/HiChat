package dao

import (
	"HiChat/global"
	"HiChat/models"
	"errors"
	"go.uber.org/zap"
)

func FriendList(userId uint) (*[]models.UserBasic, error) {
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

func AddFriend(userId, TargetId uint) (int, error) {
	if userId == TargetId {
		return -2, errors.New("不能添加自己为好友")
	}
	targetUser, err := FindUserID(TargetId)
	if err != nil {
		return -1, errors.New("未查询到用户")
	}
	if targetUser.ID == 0 {
		return -1, errors.New("未查询到用户")
	}

	relation := models.Relation{}

	if tx := global.DB.Where(" owner_id = ? and target_id =? and type =1 ", userId, TargetId).First(&relation); tx.RowsAffected == 1 {
		zap.S().Info("该好友已存在")
		return 0, errors.New("该好友已存在")
	}

	shiwu := global.DB.Begin()
	relation.OwnerId = userId
	relation.TargetID = TargetId
	relation.Type = 1

	if tx := global.DB.Create(&relation); tx.RowsAffected == 0 {
		zap.S().Info("创建失败")

		//事务回滚
		shiwu.Rollback()
		return -1, errors.New("创建好友记录失败")
	}
	shiwu.Commit()
	return 1, nil
}

func AddFriendByName(userId uint, targetName string) (int, error) {
	targetUser, err := FindUserByName(targetName)
	if err != nil {
		return -1, errors.New("未查询到用户")
	}
	if targetUser.ID == 0 {
		return -1, errors.New("未查询到用户")
	}
	return AddFriend(userId, targetUser.ID)
}
