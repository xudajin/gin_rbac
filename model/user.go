package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Name     string `grom:"not null" json:"name"`
	Password string `gorm:"not null" json:"password"`
	RoleID   uint64 `json:"role_id"`
}

// 通过姓名查询用户是否存在
func IsExistUserByName(name string) (bool, error) {
	user := User{}
	err := DB.Where("name=?", name).Find(&user).Error
	if err != nil {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

//创建用户
func AddUser(user *User) error {
	user.RoleID = 1
	if err := DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//查询单个的用户
func QueryUserById(userId uint64) (*User, error) {
	user := User{}
	err := DB.Table("users").Select("users.id,users.name,users.role_id").Where("id= ?", userId).Find(&user).Error
	if err != nil {
		return nil, err
	}
	if user.ID > 0 {

		return &user, nil
	}
	return nil, nil
}

//修改用户
func UPdateById(userId uint64, data *User) error {
	err := DB.Model(data).Omit("role_id").Where("id=?", userId).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

// 删除用户
func DeleteById(userId uint64) error {
	var user = User{}
	err := DB.Where("id=?", userId).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// 钩子
func (u *User) BeforeCreate(scope *gorm.Scope) {
	var err error
	// 密码加密
	newPwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return
	}
	u.Password = string(newPwd)
	return

}
