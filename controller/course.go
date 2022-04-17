// @Title : course
// @Description :课程接口层
// @Author : MX
// @Update : 2022/4/16 20:58

package controller

import (
	"net/http"

	"CourseSeletionSystem/service"
	"CourseSeletionSystem/util"
	"github.com/gin-gonic/gin"
)

func CreateCourse(c *gin.Context) {
	err := service.CreateCourse(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "课程信息创建", err)
	} else {
		util.SuccessResp(c, "课程信息创建")
	}
	return
}

func UpdateCourse(c *gin.Context) {
	err := service.UpdateCourse(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "课程信息修改", err)
	} else {
		util.SuccessResp(c, "课程信息修改")
	}
	return
}

func GetCourse(c *gin.Context) {
	CourseResp, err := service.GetCourse(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "获取", err)
	} else {
		util.SuccessResp(c, "获取", CourseResp)
	}
}

func DeleteCourse(c *gin.Context) {
	err := service.DeleteCourse(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "删除", err)
	} else {
		util.SuccessResp(c, "删除")
	}
}
