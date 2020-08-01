package controller

import (
	"go_web/model"
	"go_web/service"
	"go_web/util"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// 获取权限列表:分页
func PermissionList(c *gin.Context) {
	ps := service.PermissionService{}
	pageNum, _ := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 64)
	list := ps.List(pageNum)
	if list == nil {
		util.Response(c, http.StatusBadRequest, 400, "查询错误", "")
	}
	util.Response(c, http.StatusOK, 200, "查询成功", list)
}

// 添加权限
func AddPermission(c *gin.Context) {
	var data = model.Permission{}
	if err := c.BindJSON(&data); err != nil {
		util.Response(c, http.StatusBadRequest, 400, "参数错误", "")
		return
	}
	valid := validation.Validation{}
	valid.Required(data.Name, "name").Message("权限名称必须填")
	if valid.HasErrors() {
		util.Response(c, http.StatusBadRequest, 400, valid.Errors[0].Message, "")
		return
	}
	ps := service.PermissionService{}
	isExist, err := ps.Check(data.Name)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 400, "数据库错误", "")
		return
	}
	if isExist {
		util.Response(c, http.StatusBadRequest, 400, "权限已存在", "")
		return
	}
	if err = ps.Add(&data); err != nil {
		util.Response(c, http.StatusBadRequest, 400, "创建权限错误", "")
	}
	util.Response(c, http.StatusOK, 201, "创建成功", "")
}

// 修改权限
func UpdatePermission(c *gin.Context) {
	permissionID, err := strconv.ParseUint(c.Param("permission_id"), 10, 64)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 400, "参数错误", "")
	}
	var data = model.Permission{}
	if err := c.BindJSON(&data); err != nil {
		util.Response(c, http.StatusBadRequest, 400, "参数错误", "")
		return
	}
	ps := service.PermissionService{}
	if !(ps.Update(permissionID, &data)) {
		util.Response(c, http.StatusBadRequest, 400, "修改权限错误", "")
	}
	util.Response(c, http.StatusOK, 200, "修改权限成功", "")
}

// 删除权限
func DeletePermission(c *gin.Context) {
	permissionID, err := strconv.ParseUint(c.Param("permission_id"), 10, 64)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 400, "参数错误", "")
		return
	}
	ps := service.PermissionService{}
	if !(ps.Delete(permissionID)) {
		util.Response(c, http.StatusBadRequest, 400, "删除权限失败", "")
		return
	}
	util.Response(c, http.StatusOK, 200, "删除权限成功", "")
}
