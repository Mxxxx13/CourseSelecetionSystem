// @Title : teacher
// @Description :教师接口层
// @Author : MX
// @Update : 2022/4/17 14:40

package controller

import (
	"net/http"

	"CourseSeletionSystem/service"
	"CourseSeletionSystem/util"
	"github.com/gin-gonic/gin"
)

func TeacherRegister(c *gin.Context) {
	err := service.Register(c, "teacher")
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "注册", err)
	} else {
		util.SuccessResp(c, "注册")
	}
	return
}

func CreateTeacher(c *gin.Context) {
	err := service.CreateTeacher(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "教师信息创建", err)
	} else {
		util.SuccessResp(c, "教师信息创建")
	}
	return
}

func UpdateTeacher(c *gin.Context) {
	err := service.UpdateTeacher(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "教师信息修改", err)
	} else {
		util.SuccessResp(c, "教师信息修改")
	}
	return
}

func GetTeacher(c *gin.Context) {
	teacherResp, err := service.GetTeacher(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "获取", err)
	} else {
		util.SuccessResp(c, "获取", teacherResp)
	}
}

func DeleteTeacher(c *gin.Context) {
	err := service.DeleteTeacher(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "删除", err)
	} else {
		util.SuccessResp(c, "删除")
	}
}

func GetStudentSelection(c *gin.Context) {
	resp, err := service.GetStudentSelection(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "获取", err)
	} else {
		util.SuccessResp(c, "获取", resp)
	}
}
