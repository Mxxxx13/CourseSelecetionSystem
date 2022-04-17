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

func CreateStudent(c *gin.Context) {
	err := service.CreateStudent(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "学生信息创建", err)
	} else {
		util.SuccessResp(c, "学生信息创建")
	}
	return
}

func UpdateStudent(c *gin.Context) {
	err := service.UpdateStudent(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "学生信息修改", err)
	} else {
		util.SuccessResp(c, "学生信息修改")
	}
	return
}

func GetStudent(c *gin.Context) {
	studentResp, err := service.GetStudent(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "获取", err)
	} else {
		util.SuccessResp(c, "获取", studentResp)
	}
}

func DeleteStudent(c *gin.Context) {
	err := service.DeleteStudent(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "删除", err)
	} else {
		util.SuccessResp(c, "删除")
	}
}
