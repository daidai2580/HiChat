package dao

import (
	"HiChat/global"
	"HiChat/models"
	"errors"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func GetUserList() ([]*models.UserBasic, error) {
	var list []*models.UserBasic
	tx := global.DB.Find(&list)
	if tx.RowsAffected == 0 {
		return nil, errors.New("获取用户列表失败")
	}
	return list, nil
}

func FindUserByNameAndPwd(name, password string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	tx := global.DB.Where("name = ? and password=? ", name, password).First(&user)
	if tx.RowsAffected == 0 {
		return nil, errors.New("为查询到该用户")
	}

	//登录识别
	t := strconv.Itoa(int(time.Now().Unix()))

	tx1 := global.DB.Model(&user).Where("id =?", user.ID).Update("identity", t)
	if tx1.RowsAffected == 0 {
		return nil, errors.New("写入identity失败")
	}
	return &user, nil
}

func FindUserByName(name string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where("name =?", name).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("没有查询到")
	}
	return &user, nil
}

func FindUser(name string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where("name =?", name); tx.RowsAffected == 0 {
		return nil, errors.New("当前用户名已存在")
	}
	return &user, nil
}

func FindUserID(ID uint) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where(ID).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("为查询到记录")
	}
	return &user, nil
}

func FindUserByPhone(phone string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where("phone =?", phone).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("为查询到记录")
	}
	return &user, nil
}

func FindUserByEmail(email string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where("email =?", email).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("为查询到记录")
	}
	return &user, nil
}

func CreateUser(user models.UserBasic) (*models.UserBasic, error) {
	if tx := global.DB.Create(&user); tx.RowsAffected == 0 {
		zap.S().Info("新建用户数失败")
		return nil, errors.New("新增用户失败")
	}
	return &user, nil
}

func UpdateUser(user models.UserBasic) (*models.UserBasic, error) {
	tx := global.DB.Model(&user).Updates(models.UserBasic{
		Name:     user.Name,
		Password: user.Password,
		Gender:   user.Gender,
		Phone:    user.Phone,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Salt:     user.Salt,
	})
	if tx.RowsAffected == 0 {
		zap.S().Info("更新用户失败")
		return nil, errors.New("更新用户失败")
	}
	return &user, nil
}

func DeleteUser(user models.UserBasic) error {
	if tx := global.DB.Delete(&user); tx.RowsAffected == 0 {
		zap.S().Info("删除用户失败")
		return errors.New("删除用户失败")
	}
	return nil
}
