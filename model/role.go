package model

import "fmt"

type Role struct {
	BaseModel
	Name        string        `gorm:"unique;not null"`
	Permissions []*Permission `gorm:"many2many:role_permission" json:"permission"`
}

// 通过姓名查询角色是否存在
func IsExistRoleByName(name string) (bool, error) {
	role := Role{}
	err := DB.Where("name=?", name).Find(&role).Error
	if err != nil {
		return false, err
	}
	if role.ID > 0 {
		return true, nil
	}
	return false, nil
}

//创建角色
func AddRole(role *Role) error {
	if err := DB.Create(role).Error; err != nil {
		return err
	}
	return nil
}

// 查询角色
func QueryRoles() ([]*Role, error) {
	roleList := make([]*Role, 10)
	if err := DB.Select("id,name").Find(&roleList).Error; err != nil {
		return roleList, err
	}
	return roleList, nil
}

// 修改角色
func UpdateRole(roleID uint64, data *Role) error {
	if err := DB.Model(data).Select("name").Where("id=?", roleID).Update(data).Error; err != nil {
		return err
	}
	return nil
}

// 删除角色
func DeleteRole(roleID uint64) bool {
	role := Role{}
	if err := DB.Where("id=?", roleID).Delete(&role).Error; err != nil {
		return false
	}
	return true
}

// 关联权限
func RoleAddPermission(roleID uint64, permissionsID []uint64) bool {
	role := Role{}
	err := DB.Where("id=?", roleID).Find(&role).Error
	if err != nil {
		return false
	}
	// 接收查询到的权限对象
	permissions := make([]*Permission, 10)
	wrong := DB.Where("id in (?)", permissionsID).Find(&permissions).Error
	if wrong != nil {
		return false
	}
	linkErr := DB.Model(&role).Association("permissions").Append(&permissions).Error
	if linkErr != nil {
		return false
	}
	return true
}

// 通过用户名称查权限
func QueryPermissionByUserName(name string) (*Role, error) {
	role := Role{}
	// 获取用户关联角色
	err := DB.Table("users").Select("roles.id,roles.name").Joins("left join roles on roles.id = users.role_id").Where("users.name=?", name).Find(&role).Error
	if err != nil {
		return nil, err
	}
	// 关联查询权限信息，并赋值给role对象
	err = DB.Preload("Permissions").Find(&role).Error
	if err != nil {
		fmt.Println(err)
	}
	return &role, nil
}
