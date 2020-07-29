package controller

import (
	"go_web/model"
	"go_web/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Name     string
	Password string
}

func LoginController(c *gin.Context) {
	login := Login{}
	if err := c.BindJSON(&login); err != nil {
		util.Response(c, http.StatusOK, 400, "参数错误", "")
		return
	}

	user := model.User{}
	if err := model.DB.Where("name=? AND password=?", login.Name, login.Password).Find(&user).Error; err != nil {
		util.Response(c, http.StatusBadRequest, 400, "用户名或密码错误", "")
		return
	}
	if user.ID > 0 {
		token, err := util.GenerateToken(user.Name, user.Password)
		if err != nil {
			util.Response(c, http.StatusBadRequest, 400, "签发Token错误", "")
		}
		util.Response(c, http.StatusOK, 200, "签发Token成功", token)
	}

}
