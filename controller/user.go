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
	id, err := service.Login(c)
	username := c.PostForm("username")
	role := c.PostForm("role")
	if err != nil {
		msg := "用户名或密码错误"
		util.ErrorResp(c, http.StatusBadRequest, msg, err)

	} else {
		jwt := util.NewJWT(id, username, role)
		msg := "欢迎回来" + username

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": msg,
			"token":   jwt.Token,
		})
	}
}

// AlterUser
func AlterUser(c *gin.Context) {
	err := service.AlterUser(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "修改失败,用户名不存在", err)
	} else {
		util.SuccessResp(c, "修改成功")
	}
}

// TODO: 添加修改密码的功能
