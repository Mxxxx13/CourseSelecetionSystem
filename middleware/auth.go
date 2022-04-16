// @Title : auth
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/16 19:36

package middleware

import (
	"net/http"

	"CourseSeletionSystem/util"
	"github.com/gin-gonic/gin"
)

func LoginRequired(c *gin.Context) {
	token := c.PostForm("token")
	//根据token解析出jwt
	jwt, err := util.CheckJWT(token)
	if err != nil {
		util.ErrorResp(c, http.StatusUnauthorized, "需要登录", err)
		c.Abort()
		return
	}

	// 设置id方便后续操作
	c.Set("uid", jwt.Payload.Sub.Uid)
	c.Set("role", jwt.Payload.Sub.Role)
	c.Next()
}
