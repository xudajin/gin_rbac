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

// 添加用户
func AddUser(c *gin.Context) {
	var data = model.User{}
	if err := c.BindJSON(&data); err != nil {
		util.Response(c, http.StatusBadRequest, 400, "参数错误", "")
		return
	}
	valid := validation.Validation{}
	valid.Required(data.Name, "name").Message("姓名必须填")
	valid.Required(data.Password, "password").Message("密码必须填")
	valid.MinSize(data.Name, 1, "name").Message("名字太短")
	valid.MinSize(data.Password, 6, "password").Message("密码必须大于6位数")
	if valid.HasErrors() {
		util.Response(c, http.StatusBadRequest, 400, valid.Errors[0].Message, "")
		return
	}
	us := service.UserService{}
	isExist, err := us.Check(data.Name)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 400, "数据库错误", "")
		return
	}
	if isExist {
		util.Response(c, http.StatusBadRequest, 400, "用户已存在", "")
		return
	}
	if err = us.Add(&data); err != nil {
		util.Response(c, http.StatusBadRequest, 400, "创建用户错误", "")
		return
	}
	util.Response(c, http.StatusOK, 201, "创建成功", "")

}

// 查询用户列表
func QueryUserList(c *gin.Context) {
	us := service.UserService{}
	pageNum, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 64)
	userList, err := us.QueryUserList(pageNum)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 400, "查询错误", "")
		return
	}
	util.Response(c, http.StatusOK, 200, "查询成功", userList)
}

// 修改用户
func UpdateUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 400, "参数错误", "")
		return
	}
	data := model.User{}
	if err := c.BindJSON(&data); err != nil {
		util.Response(c, http.StatusBadRequest, 400, "参数错误", "")
		return
	}
	valid := validation.Validation{}
	if data.Name != "" {
		valid.MinSize(data.Name, 1, "name").Message("名字太短")
	}
	if data.Password != "" {
		valid.MinSize(data.Name, 6, "name").Message("密码必须大于6位数")
	}
	if valid.HasErrors() {
		util.Response(c, http.StatusBadRequest, 400, valid.Errors[0].Message, "")
		return
	}
	us := service.UserService{}
	err = us.UpdateByID(userID, &data)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 400, "修改用户错误", "")
		return
	}
	util.Response(c, http.StatusOK, 200, "修改用户成功", "")
}

// 修改用户密码
func ChangePassword(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 400, "参数错误", "")
		return
	}
	data := model.User{}
	if err := c.BindJSON(&data); err != nil {
		util.Response(c, http.StatusBadRequest, 400, "参数错误", "")
		return
	}
	valid := validation.Validation{}
	valid.Required(data.Password, "password").Message("密码必须填")
	valid.MinSize(data.Password, 6, "password").Message("密码不能小于6位数")
	if valid.HasErrors() {
		util.Response(c, http.StatusBadRequest, 400, valid.Errors[0].Message, "")
		return
	}
	us := service.UserService{}
	if !(us.ChangePassword(userID, &data)) {
		util.Response(c, http.StatusBadRequest, 400, "修改密码错误", "")
		return
	}
	util.Response(c, http.StatusOK, 200, "修改密码成功", "")

}

// 删除用户
func DeleteUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		util.Response(c, http.StatusBadRequest, 400, "参数错误", "")
		return
	}
	us := service.UserService{}
	if err := us.DeleteByID(userID); err != nil {
		util.Response(c, http.StatusBadRequest, 400, "删除用户错误", "")
	}

}
