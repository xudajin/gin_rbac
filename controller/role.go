package controller

import (
	"encoding/json"
	e "go_web/error"
	"go_web/model"
	"go_web/service"
	"go_web/util"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// 添加角色
func AddRole(c *gin.Context) {
	var data = model.Role{}
	if err := c.BindJSON(&data); err != nil {
		util.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, e.Msg(e.INVALID_PARAMS), "")
		return
	}
	valid := validation.Validation{}
	valid.Required(data.Name, "name").Message("角色名称必须填")
	if valid.HasErrors() {
		util.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors[0].Message, "")
		return
	}
	rs := service.RoleService{}
	isExist, err := rs.Check(data.Name)
	if err != nil {
		util.Response(c, http.StatusBadRequest, e.ERROR, e.Msg(e.ERROR), "")
		return
	}
	if isExist {
		util.Response(c, http.StatusBadRequest, e.ERROR, e.Msg(e.ERROR), "")
		return
	}
	if err = rs.AddRole(&data); err != nil {
		util.Response(c, http.StatusBadRequest, e.ERROR, e.Msg(e.ERROR), "")
		return
	}
	util.Response(c, http.StatusOK, e.SUCCESS, e.Msg(e.SUCCESS), "")

}

// 查询角色列表
func QueryRoles(c *gin.Context) {
	rs := service.RoleService{}
	roleList, err := rs.QueryRoles()
	if err != nil {
		util.Response(c, http.StatusBadRequest, e.ERROR, e.Msg(e.ERROR), "")
		return
	}
	count := len(roleList)
	util.ListResponse(c, http.StatusBadRequest, e.SUCCESS, e.Msg(e.SUCCESS), count, roleList)
}

// 修改角色信息
func UpdateRole(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("role_id"), 10, 64)
	if err != nil {
		util.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, e.Msg(e.INVALID_PARAMS), "")
		return
	}
	data := model.Role{}
	if err = c.BindJSON(&data); err != nil {
		util.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, e.Msg(e.INVALID_PARAMS), "")
		return
	}
	// 验证参数
	valid := validation.Validation{}
	valid.Required(data.Name, "name").Message("角色名称必须填")
	valid.MinSize(data.Name, 2, "name").Message("角色名称必不能小于2个字")
	if valid.HasErrors() {
		util.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors[0].Message, "")
		return
	}

	rs := service.RoleService{}
	err = rs.UpdateRole(roleID, &data)
	if err != nil {
		util.Response(c, http.StatusBadRequest, e.ERROR, e.Msg(e.ERROR), "")
		return
	}
	util.Response(c, http.StatusOK, e.SUCCESS, e.Msg(e.SUCCESS), "")
}

// 删除角色
func DeleteRole(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("role_id"), 10, 64)
	if err != nil {
		util.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, e.Msg(e.INVALID_PARAMS), "")
		return
	}
	rs := service.RoleService{}
	if !(rs.DeleteRole(roleID)) {
		util.Response(c, http.StatusBadRequest, e.ERROR, e.Msg(e.ERROR), "")
		return
	}
	util.Response(c, http.StatusOK, e.SUCCESS, e.Msg(e.SUCCESS), "")
}

// 角色权限关联
func RoleAddPermission(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("role_id"), 10, 64)
	if err != nil {
		util.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, e.Msg(e.INVALID_PARAMS), "")
	}
	// 读取post.body中的数据
	body, _ := ioutil.ReadAll(c.Request.Body)
	if len(body) == 0 {
		util.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, e.Msg(e.INVALID_PARAMS), "")
		return
	}
	// 定义map接收
	var permissionID map[string][]uint64
	//将body的数据转成json,并赋值给 permissionID
	json.Unmarshal(body, &permissionID)

	rs := service.RoleService{}
	if !(rs.RoleAddPermission(roleID, permissionID["permission_id"])) {
		util.Response(c, http.StatusBadRequest, e.ERROR, e.Msg(e.ERROR), "")
		return
	}
	util.Response(c, http.StatusOK, e.SUCCESS, e.Msg(e.SUCCESS), "")
}

// 通过角色id获取权限
func QueryPermissionByRoleID(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("role_id"), 10, 64)
	if err != nil {
		util.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, e.Msg(e.INVALID_PARAMS), "")
		return
	}
	rs := service.RoleService{}
	permissionList, ok := rs.GetRolePermissionByID(roleID)
	if !ok {
		util.Response(c, http.StatusBadRequest, e.ERROR, e.Msg(e.ERROR), "")
		return
	}
	count := len(permissionList) // 获取返回数据个数
	util.ListResponse(c, http.StatusOK, e.SUCCESS, e.Msg(e.SUCCESS), count, permissionList)

}
