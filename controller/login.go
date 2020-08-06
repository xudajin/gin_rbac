package controller

import (
	"go_web/model"
	"go_web/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func LoginController(c *gin.Context) {
	login := Login{}
	if err := c.BindJSON(&login); err != nil {
		util.Response(c, http.StatusOK, 400, "参数错误", "")
		return
	}
	user := model.User{}
	if err := model.DB.Where("name=?", login.Name).Find(&user).Error; err != nil {
		util.Response(c, http.StatusBadRequest, 400, "用户名错误", "")
		return
	}
	// 验证密码
	if user.ID > 0 {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
		if err != nil {
			util.Response(c, http.StatusBadRequest, 400, "密码错误", "")
			return
		}
		// 签发Token
		token, err := util.GenerateToken(user.ID, user.Name)
		if err != nil {
			util.Response(c, http.StatusBadRequest, 400, "签发Token错误", "")
			return
		}
		util.Response(c, http.StatusOK, 200, "签发Token成功", token)
		return
	}
	util.Response(c, http.StatusBadRequest, 400, "用户不存在", "")
}
