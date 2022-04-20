// @Title : studentCheck
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/20 16:33

package middleware

import (
	"errors"
	"net/http"

	"CourseSeletionSystem/util"
	"github.com/gin-gonic/gin"
)

func StudentCheck(c *gin.Context) {
	rolestr, exists := c.Get("role")
	if !exists {
		util.ErrorResp(c, http.StatusBadRequest, "失败", errors.New("role not exists"))
		c.Abort()
	}
	role := rolestr.(string)
	if role != "student" {
		util.ErrorResp(c, http.StatusBadRequest, "失败", errors.New("权限不够"))
		c.Abort()
	}

	c.Next()
}

func TeacherCheck(c *gin.Context) {
	rolestr, exists := c.Get("role")
	if !exists {
		util.ErrorResp(c, http.StatusBadRequest, "失败", errors.New("role not exists"))
		c.Abort()
	}
	role := rolestr.(string)
	if role != "teacher" {
		util.ErrorResp(c, http.StatusBadRequest, "失败", errors.New("权限不够"))
		c.Abort()
	}

	c.Next()
}

func AdminCheck(c *gin.Context) {
	rolestr, exists := c.Get("role")
	if !exists {
		util.ErrorResp(c, http.StatusBadRequest, "失败", errors.New("role not exists"))
		c.Abort()
	}
	role := rolestr.(string)
	if role != "admin" {
		util.ErrorResp(c, http.StatusBadRequest, "失败", errors.New("权限不够"))
		c.Abort()
	}

	c.Next()
}

func AdminAndTeacherCheck(c *gin.Context) {
	rolestr, exists := c.Get("role")
	if !exists {
		util.ErrorResp(c, http.StatusBadRequest, "失败", errors.New("role not exists"))
		c.Abort()
	}
	role := rolestr.(string)
	if role == "admin" || role == "teacher" {
		c.Next()
	}

	util.ErrorResp(c, http.StatusBadRequest, "失败", errors.New("权限不够"))
	c.Abort()

}

func AdminAndStudentCheck(c *gin.Context) {
	rolestr, exists := c.Get("role")
	if !exists {
		util.ErrorResp(c, http.StatusBadRequest, "失败", errors.New("role not exists"))
		c.Abort()
	}
	role := rolestr.(string)
	if role == "admin" || role == "teacher" {
		c.Next()
	}

	util.ErrorResp(c, http.StatusBadRequest, "失败", errors.New("权限不够"))
	c.Abort()

}
