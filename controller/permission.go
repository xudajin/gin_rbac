package controller

import (
	e "go_web/error"
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
		util.Response(c, http.StatusBadRequest, e.ERROR, e.Msg(e.ERROR), "")
	}
	count := len(list)
	util.ListResponse(c, http.StatusOK, e.SUCCESS, e.Msg(e.SUCCESS), count, list)
}

// 添加权限
func AddPermission(c *gin.Context) {
	var data = model.Permission{}
	if err := c.BindJSON(&data); err != nil {
		util.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, e.Msg(e.INVALID_PARAMS), "")
		return
	}
	valid := validation.Validation{}
	valid.Required(data.Name, "name").Message("权限名称必须填")
	if valid.HasErrors() {
		util.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors[0].Message, "")
		return
	}
	ps := service.PermissionService{}
	isExist, err := ps.Check(data.Name)
	if err != nil {
		util.Response(c, http.StatusBadRequest, e.ERROR, e.Msg(e.ERROR), "")
		return
	}
	if isExist {
		util.Response(c, http.StatusBadRequest, e.PERMISSION_ISEXIST, e.Msg(e.PERMISSION_ISEXIST), "")
		return
	}
	if err = ps.Add(&data); err != nil {
		util.Response(c, http.StatusBadRequest, e.ERROR, e.Msg(e.ERROR), "")
	}
	util.Response(c, http.StatusOK, e.SUCCESS, e.Msg(e.SUCCESS), "")
}

// 修改权限
func UpdatePermission(c *gin.Context) {
	permissionID, err := strconv.ParseUint(c.Param("permission_id"), 10, 64)
	if err != nil {
		util.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, e.Msg(e.INVALID_PARAMS), "")
	}
	var data = model.Permission{}
	if err := c.BindJSON(&data); err != nil {
		util.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, e.Msg(e.INVALID_PARAMS), "")
		return
	}
	ps := service.PermissionService{}
	if !(ps.Update(permissionID, &data)) {
		util.Response(c, http.StatusBadRequest, e.ERROR, e.Msg(e.ERROR), "")
	}
	util.Response(c, http.StatusOK, e.SUCCESS, e.Msg(e.SUCCESS), "")
}

// 删除权限
func DeletePermission(c *gin.Context) {
	permissionID, err := strconv.ParseUint(c.Param("permission_id"), 10, 64)
	if err != nil {
		util.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, e.Msg(e.INVALID_PARAMS), "")
		return
	}
	ps := service.PermissionService{}
	if !(ps.Delete(permissionID)) {
		util.Response(c, http.StatusBadRequest, e.ERROR, e.Msg(e.ERROR), "")
		return
	}
	util.Response(c, http.StatusOK, e.SUCCESS, e.Msg(e.SUCCESS), "")
}
