package dao

import (
	"context"

	"gorm.io/gorm"

	// "serein/pkg/utils/log"
	"serein/repository/db/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// ExistOrNotByEmail 根据Email判断是否存在该名字
func (dao *UserDao) ExistOrNotByEmail(email string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	if count == 0 {
		return user, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

// GeuUserByID 根据user_id获取用户相关信息
func (dao *UserDao) GeuUserByID(uID uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).
		Where("id = ?", uID).
		First(&user).Error
	return
}

func (dao *UserDao) UpdateUserById(uID uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id=?", uID).
		Updates(&user).Error
}
