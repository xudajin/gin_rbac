package service

import (
	"go_web/model"
	"go_web/serializer"

	"github.com/jinzhu/gorm"
)

type RoleService struct {
}

// 检查权限是否存在
func (rs *RoleService) Check(name string) (bool, error) {
	isExist, err := model.IsExistRoleByName(name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if isExist {
		return true, nil
	}
	return false, nil
}

// 添加角色
func (rs *RoleService) AddRole(role *model.Role) error {
	if err := model.AddRole(role); err != nil {
		return err
	}
	return nil
}

// 查看角色列表
func (rs *RoleService) QueryRoles() ([]*model.Role, error) {
	roleList, err := model.QueryRoles()

	if err != nil {
		return roleList, err
	}
	return roleList, nil
}

// 修改角色
func (rs *RoleService) UpdateRole(roleID uint64, data *model.Role) error {
	err := model.UpdateRole(roleID, data)
	if err != nil {
		return err
	}
	return nil
}

// 删除角色
func (rs *RoleService) DeleteRole(roleID uint64) bool {
	if !(model.DeleteRole(roleID)) {
		return false
	}
	return true
}

// 角色添加权限
func (rs *RoleService) RoleAddPermission(roleID uint64, permissionsID []uint64) bool {
	if !(model.RoleAddPermission(roleID, permissionsID)) {
		return false
	}
	return true
}

// 通过角色id获取拥有权限
func (rs *RoleService) GetRolePermissionByID(roleID uint64) ([]*serializer.TreePermission, bool) {
	TreePermissionList, ok := model.QueryPermissionsByRoleID(roleID)
	if !ok {
		return nil, false
	}
	// 对查询数据进行结构化处理.形成树状结构
	permissionList := []*serializer.TreePermission{}
	for _, permission := range TreePermissionList {
		for _, cpermission := range TreePermissionList {
			if uint64(permission.ID) == cpermission.ParentID {
				permission.Childen = append(permission.Childen, cpermission)
			}
		}
		// 让该权限是root权限时，才添加入权限列表
		if permission.Category == "root" {
			permissionList = append(permissionList, permission)
		}
	}
	return permissionList, true
}
