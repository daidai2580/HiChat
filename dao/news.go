package dao

import (
	"HiChat/global"
	"HiChat/models"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

type JSON json.RawMessage

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

func GetNewsAndUser2() (j JSON, e error) {

	var s JSON
	tx := global.DB.Table("news").Select("news.title,news.content,news.author_id,user_basics.name").Joins("left join user_basics on news.author_id= user_basics.id").Scan(&s)
	if tx.RowsAffected == 0 {
		return nil, errors.New("获取新闻列表失败")
	}
	return s, nil
}

func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = JSON("null")
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		if len(v) > 0 {
			bytes = make([]byte, len(v))
			copy(bytes, v)
		}
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage(bytes)
	*j = JSON(result)
	return nil
}
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return string(j), nil
}
