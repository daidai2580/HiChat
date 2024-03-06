package dao

import (
	"HiChat/global"
	"HiChat/models"
	"errors"
	"go.uber.org/zap"
)

func GetNewsList() ([]*models.News, error) {
	var list []*models.News
	tx := global.DB.Find(&list)
	if tx.RowsAffected == 0 {
		return nil, errors.New("获取新闻列表失败")
	}
	return list, nil
}

func AddNews(news models.News) (*models.News, error) {
	tx := global.DB.Create(&news)
	if tx.RowsAffected == 0 {
		zap.S().Info("新建新闻失败")
		return nil, errors.New("新建新闻失败")
	}
	return &news, nil
}

func GetNewsAndUser() ([]*models.NewsAndUser, error) {

	var list []*models.NewsAndUser
	tx := global.DB.Table("news").Select("news.title,news.content,news.author_id,user_basics.name").Joins("left join user_basics on news.author_id= user_basics.id").Find(&list)
	if tx.RowsAffected == 0 {
		return nil, errors.New("获取新闻列表失败")
	}
	return list, nil
}
