package model

type Permission struct {
	BaseModel
	Name   string  `gorm:"not null;unique" json:"name"`
	Path   string  `json:"path"`
	Method string  `json:"method"`
	Roles  []*Role `gorm:"many2many:role_permission" json:"role"`
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
func PermissionList() (*[]Permission, bool) {
	permissionList := []Permission{}
	if err := DB.Find(&permissionList).Error; err != nil {
		return &permissionList, false
	}
	return &permissionList, true
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
