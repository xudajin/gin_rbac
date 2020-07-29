package service

import (
	"go_web/model"

	"github.com/jinzhu/gorm"
)

type PermissionService struct {
}

// 检查权限是否重复
func (ps *PermissionService) Check(name string) (bool, error) {
	isExist, err := model.IsExistPermissionByName(name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if isExist {
		return true, nil
	}
	return false, nil
}

// 获取权限列表
func (ps *PermissionService) List() interface{} {
	list, ok := model.PermissionList()
	if !ok {
		return nil
	}
	return list
}

// 添加权限
func (ps *PermissionService) Add(permission *model.Permission) error {
	if err := model.AddPermission(permission); err != nil {
		return err
	}
	return nil
}

// 修改权限
func (ps *PermissionService) Update(id uint64, data *model.Permission) bool {
	if !(model.UpdatePermission(id, data)) {
		return false
	}
	return true

}
