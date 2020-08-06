package model

import "go_web/serializer"

type Role struct {
	BaseModel
	Name   string `gorm:"unique;not null"`
	Status uint64 `grom:"default(1)" json:"status"`
}

// 角色关联权限表
type RolePermission struct {
	ID           uint   `gorm:"primary_key" json:"id"`
	RoleID       uint64 `gorm:"not null" json:"role_id"`
	PermissionID uint64 `gorm:"not null" json:"permission_id"`
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

// 查询角色列表
func QueryRoles() ([]*Role, error) {
	roleList := make([]*Role, 10)
	if err := DB.Select("id,name").Find(&roleList).Error; err != nil {
		return roleList, err
	}
	return roleList, nil
}

// 修改角色
func UpdateRole(roleID uint64, data *Role) error {
	role := Role{}
	if err := DB.Model(&role).Select("name").Where("id=?", roleID).Update(data).Error; err != nil {
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
	// 清空角色关联的权限
	rolePermission := RolePermission{}
	err = DB.Where("role_id=?", role.ID).Delete(&rolePermission).Error
	if err != nil {
		return false
	}
	// 接收查询到的权限对象
	permissions := make([]*Permission, 10)
	wrong := DB.Where("id in (?)", permissionsID).Find(&permissions).Error
	if wrong != nil {
		return false
	}
	for _, v := range permissions {
		//查询一条记录，若查询不到，则创建记录 FirstOrCreate
		err = DB.Where(&RolePermission{RoleID: uint64(role.ID), PermissionID: uint64(v.ID)}).FirstOrCreate(&RolePermission{RoleID: uint64(role.ID), PermissionID: uint64(v.ID)}).Error
		if err != nil {
			return false
		}
	}
	return true
}

// 通过角色id获取权限
func QueryPermissionsByRoleID(roleID uint64) ([]*serializer.TreePermission, bool) {
	role := Role{}
	err := DB.Select("id,name").Where("id=?", roleID).Find(&role).Error
	if err != nil {
		return nil, false
	}
	PermissionIDList := []*RolePermission{}
	err = DB.Select("permission_id").Where("role_id=?", role.ID).Find(&PermissionIDList).Error
	if err != nil {
		return nil, false
	}
	IDList := []uint64{}
	for _, v := range PermissionIDList {
		IDList = append(IDList, v.PermissionID)
	}
	permissionList := []*serializer.TreePermission{}
	// 原生sql join 查询
	err = DB.Table("permissions").Select("id,name,path,method,category,parent_id").Where("id in (?) AND deleted_at is null", IDList).Scan(&permissionList).Error
	if err != nil {
		return permissionList, false
	}
	return permissionList, true
}
