package controller

import (
	e "go_web/error"
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
		util.Response(c, http.StatusOK, e.INVALID_PARAMS, e.Msg(e.INVALID_PARAMS), "")
		return
	}
	user := model.User{}
	if err := model.DB.Where("name=?", login.Name).Find(&user).Error; err != nil {
		util.Response(c, http.StatusBadRequest, e.USER_NOT_FIND, e.Msg(e.USER_NOT_FIND), "")
		return
	}
	// 验证密码
	if user.ID > 0 {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
		if err != nil {
			util.Response(c, http.StatusBadRequest, e.PASSWORD_WRONG, e.Msg(e.PASSWORD_WRONG), "")
			return
		}
		// 签发Token
		token, err := util.GenerateToken(user.ID, user.Name, user.RoleID)
		if err != nil {
			util.Response(c, http.StatusBadRequest, e.ERROR, e.Msg(e.ERROR), "")
			return
		}
		util.Response(c, http.StatusOK, e.SUCCESS, e.Msg(e.SUCCESS), token)
		return
	}
}
