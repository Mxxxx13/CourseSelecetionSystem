// @Title : controller
// @Description :学生接口层
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
		util.ErrorResp(c, http.StatusBadRequest, "注册", err)
	} else {
		util.SuccessResp(c, "注册")
	}
	return
}

func CreateStudentInfo(c *gin.Context) {
	err := service.CreateStudentInfo(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "学生信息创建", err)
	} else {
		util.SuccessResp(c, "学生信息创建")
	}
	return
}

func UpdateStudentInfo(c *gin.Context) {
	err := service.UpdateStudentInfo(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "学生信息修改", err)
	} else {
		util.SuccessResp(c, "学生信息修改")
	}
	return
}
