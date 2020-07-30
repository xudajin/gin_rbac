package service

import (
	"go_web/model"

	"github.com/jinzhu/gorm"
)

type UserService struct {
}

// 检查用户是否重复
func (us *UserService) Check(name string) (bool, error) {
	isExist, err := model.IsExistUserByName(name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if isExist {
		return true, nil
	}
	return false, nil
}

// 添加用户
func (us *UserService) Add(user *model.User) error {
	if err := model.AddUser(user); err != nil {
		return err
	}
	return nil
}

// 查询用户
func (us *UserService) QueryByID(userID uint64) (*model.User, error) {
	data, err := model.QueryUserById(userID)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 修改用户信息
func (us *UserService) UpdateByID(userID uint64, data *model.User) error {
	err := model.UPdateById(userID, data)
	if err != nil {
		return err
	}
	return nil
}

// 删除用户
func (us *UserService) DeleteById(userId uint64) error {
	err := model.DeleteById(userId)
	if err != nil {
		return err
	}
	return nil
}
