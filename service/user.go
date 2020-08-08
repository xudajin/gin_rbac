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

//查询用户列表
func (us *UserService) QueryUserList(pageNum uint64) ([]*model.User, error) {
	userList, err := model.QueryUserList(pageNum)
	if err != nil {
		return nil, err
	}
	return userList, nil

}

// 修改用户信息
func (us *UserService) UpdateByID(userID uint64, data *model.User) error {
	err := model.UpdateByID(userID, data)
	if err != nil {
		return err
	}
	return nil
}

//修改用户密码
func (us *UserService) ChangePassword(userID uint64, data *model.User) bool {
	if !(model.ChangePassword(userID, data)) {
		return false
	}
	return true
}

// 删除用户
func (us *UserService) DeleteByID(userID uint64) error {
	err := model.DeleteByID(userID)
	if err != nil {
		return err
	}
	return nil
}
