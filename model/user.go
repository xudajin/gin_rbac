package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Name     string `grom:"not null" json:"name"`
	Password string `gorm:"not null" json:"password,omitempty"`
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

// 查询用户列表
func QueryUserList(pageNum uint64) (*[]User, error) {
	userList := []User{}
	err := DB.Select("id,name,role_id").Offset((pageNum - 1) * 5).Limit(5).Find(&userList).Error
	if err != nil {
		return &userList, err
	}
	return &userList, nil
}

//修改用户
func UpdateByID(userID uint64, data *User) error {
	err := DB.Model(data).Omit("role_id,password").Where("id=?", userID).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

// 修改用户密码
func ChangePassword(userID uint64, data *User) bool {
	user := User{}
	newPwd, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return false
	}
	data.Password = string(newPwd)
	err = DB.Model(&user).Select("password").Where("id=?", userID).Updates(data).Error
	if err != nil {
		return false
	}
	return true
}

// 删除用户
func DeleteByID(userID uint64) error {
	var user = User{}
	err := DB.Where("id=?", userID).Delete(&user).Error
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
