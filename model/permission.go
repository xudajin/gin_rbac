package model

type Permission struct {
	BaseModel
	Name     string `gorm:"not null;unique" json:"name"`
	Path     string `json:"path"`
	Method   string `json:"method"`
	Code     string `gorm:"not null;unique" json:"code"`
	ParentID uint64 `json:"parent_id"`
	Remark   string `json:"remark"`
	Category string `json:"category"`
}

// 查看权限名是否存在
func IsExistPermissionByName(name string) (bool, error) {
	permission := Permission{}
	err := DB.Where("name=?", name).Find(&permission).Error
	if err != nil {
		return false, err
	}
	if permission.ID > 0 {
		return true, nil
	}
	return false, nil
}

// 获取权限列表
func PermissionList(pageNum uint64) ([]*Permission, bool) {
	permissionList := []*Permission{}
	if err := DB.Offset((pageNum - 1) * 5).Limit(5).Find(&permissionList).Error; err != nil {
		return permissionList, false
	}
	return permissionList, true
}

// 添加权限
func AddPermission(data *Permission) error {
	if err := DB.Create(data).Error; err != nil {
		return err
	}
	return nil
}

//修改权限
func UpdatePermission(id uint64, data *Permission) bool {
	permission := Permission{}
	if err := DB.Model(&permission).Where("id=? ", id).Updates(data).Error; err != nil {
		return false
	}
	return true
}

// 删除权限
func DeletePermission(id uint64) bool {
	permission := Permission{}
	err := DB.Where("id=?", id).Delete(&permission).Error
	if err != nil {
		return false
	}
	return true
}
