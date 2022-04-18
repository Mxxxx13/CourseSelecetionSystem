// @Title : courseSelection
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/17 18:37

package controller

import (
	"net/http"

	"CourseSeletionSystem/service"
	"CourseSeletionSystem/util"
	"github.com/gin-gonic/gin"
)

func StudentSelectCourse(c *gin.Context) {
	err := service.StudentSelectCourse(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "学生选课", err)
	} else {
		util.SuccessResp(c, "学生选课")
	}
	return
}

func StudentGetCourse(c *gin.Context) {
	resp, err := service.StudentGetCourse(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "查看学生选课", err)
	} else {
		util.SuccessResp(c, "查看学生选课", resp)
	}
}

func StudentDeleteCourse(c *gin.Context) {
	err := service.StudentDeleteCourse(c)
	if err != nil {
		util.ErrorResp(c, http.StatusBadRequest, "学生退课", err)
	} else {
		util.SuccessResp(c, "学生退课")
	}
}
