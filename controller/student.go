// @Title : controller
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/16 10:52

package controller

import (
	"net/http"

	"CourseSeletionSystem/service"
	"CourseSeletionSystem/util"
	"github.com/gin-gonic/gin"
)

func StudentRegister(c *gin.Context) {
	err := service.Register(c, "student")
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "注册失败，用户已存在", err)
	} else {
		util.SuccessResp(c, "注册成功")
	}
	return
}

func StudentInfo() {

}
