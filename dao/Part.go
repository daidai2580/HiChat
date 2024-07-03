package dao

import (
	"HiChat/global"
	"HiChat/models"
	"errors"
	"go.uber.org/zap"
)

func CreatePart(part models.Part) (*models.Part, error) {
	tx := global.DB.Create(&part)
	if tx.RowsAffected == 0 {
		zap.S().Error("创建物料失败")
		return nil, errors.New("创建物料失败")
	}
	return &part, nil
}

func FindPartByOid(oid string) (*models.Part, error) {
	part := models.Part{Oid: oid}
	tx := global.DB.Where(&part).Find(&part)
	if tx.RowsAffected == 0 {
		zap.S().Error("根据oid没有找到对应物料")
		return nil, errors.New("根据oid没有找到对应物料")
	}
	return &part, nil
}
