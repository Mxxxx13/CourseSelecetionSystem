// @Title : user
// @Description :用户接口层
// @Author : MX
// @Update : 2022/4/15 22:06

package controller

import (
	"net/http"

	"CourseSeletionSystem/service"
	"CourseSeletionSystem/util"
	"github.com/gin-gonic/gin"
)

//Login 返回登录接口,成功登录返回token
func Login(c *gin.Context) {
	user, err := service.Login(c)

	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "注册", err)

	} else {
		jwt := util.NewJWT(user.ID, user.Username, user.Role)
		msg := "欢迎回来" + user.Username

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": msg,
			"token":   jwt.Token,
		})
	}
}

func BindEmail(c *gin.Context) {
	err := service.BindEmail(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "绑定", err)
	} else {
		util.SuccessResp(c, "绑定")
	}
}

func AlterPassword(c *gin.Context) {
	err := service.AlterPassword(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "找回", err)
	} else {
		util.SuccessResp(c, "找回")
	}
}

func SendEmail(c *gin.Context) {
	err := service.SendEmail(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "发送", err)
	} else {
		util.SuccessResp(c, "发送")
	}
}
