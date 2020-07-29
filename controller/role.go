package controller

import (
	"encoding/json"
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
		util.Response(c, http.StatusBadRequest, 400, "参数错误", "")
		return
	}
	valid := validation.Validation{}
	valid.Required(data.Name, "name").Message("角色名称必须填")
	if valid.HasErrors() {
		util.Response(c, http.StatusBadRequest, 400, valid.Errors[0].Message, "")
		return
	}
	rs := service.RoleService{}
	isExist, err := rs.Check(data.Name)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 400, "数据库错误", "")
		return
	}
	if isExist {
		util.Response(c, http.StatusBadRequest, 400, "角色已存在", "")
		return
	}
	if err = rs.AddRole(&data); err != nil {
		util.Response(c, http.StatusBadRequest, 400, "创建角色错误", "")
		return
	}
	util.Response(c, http.StatusOK, 201, "创建成功", "")

}

// 查询角色列表
func QueryRoles(c *gin.Context) {
	rs := service.RoleService{}
	roleList, err := rs.QueryRoles()
	if err != nil {
		util.Response(c, http.StatusBadRequest, 400, "查询角色错误", "")
		return
	}
	util.Response(c, http.StatusBadRequest, 200, "查询角色成功", roleList)
}

// 修改角色信息
func UpdateRole(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("role_id"), 10, 64)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 400, "参数错误", "")
		return
	}
	data := model.Role{}
	if err = c.BindJSON(&data); err != nil {
		util.Response(c, http.StatusBadRequest, 400, "参数错误", "")
		return
	}
	// 验证参数
	valid := validation.Validation{}
	valid.Required(data.Name, "name").Message("角色名称必须填")
	valid.MinSize(data.Name, 2, "name").Message("角色名称必不能小于2个字")
	if valid.HasErrors() {
		util.Response(c, http.StatusBadRequest, 400, valid.Errors[0].Message, "")
		return
	}

	rs := service.RoleService{}
	err = rs.UpdateRole(roleID, &data)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 400, "修改角色错误", "")
		return
	}
	util.Response(c, http.StatusOK, 200, "修改角色成功", "")
}

// 删除角色
func DeleteRole(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("role_id"), 10, 64)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 400, "参数错误", "")
	}
	rs := service.RoleService{}
	if !(rs.DeleteRole(roleID)) {
		util.Response(c, http.StatusBadRequest, 400, "删除操作失败", "")
	}
	util.Response(c, http.StatusOK, 200, "删除操作成功", "")
}

// 角色权限关联
func RoleAddPermission(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("role_id"), 10, 64)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 400, "参数错误", "")
	}
	// 读取post.body中的数据
	body, _ := ioutil.ReadAll(c.Request.Body)
	// 定义map接收
	var permissionID map[string][]uint64
	//将body的数据转成json,并赋值给 permissionID
	json.Unmarshal(body, &permissionID)

	rs := service.RoleService{}
	if !(rs.RoleAddPermission(roleID, permissionID["permission_id"])) {
		util.Response(c, http.StatusBadRequest, 400, "角色权限关联错误", "")
		return
	}
	util.Response(c, http.StatusOK, 200, "角色权限关联成功", "")
}
