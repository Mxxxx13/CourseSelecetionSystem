// @Title : courseSelection
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/17 18:37

package service

import (
	"errors"
	"strconv"

	"CourseSeletionSystem/dao"
	"CourseSeletionSystem/model"
	"github.com/gin-gonic/gin"
)

func StudentSelectCourse(c *gin.Context) (err error) {
	var selection model.CourseSelection

	uid, exists := c.Get("uid")
	if !exists {
		return errors.New("uid do not exists")
	}

	student, err := dao.GetStudentByUid(uid.(uint))
	selection.StudentID = student.ID

	courseIDstr := c.PostForm("courseID")
	courseID, err := strconv.Atoi(courseIDstr)
	if err != nil {
		return
	}
	selection.CourseID = uint(courseID)

	err = dao.StudentSelectCourse(selection)
	return
}

func StudentGetCourse(c *gin.Context) (resp model.StudentCourseResp, err error) {
	uid, exists := c.Get("uid")
	if !exists {
		return resp, errors.New("uid do not exists")
	}

	student, err := dao.GetStudentByUid(uid.(uint))
	studentResp := model.StudentResp{
		Name:    student.Name,
		Number:  student.Number,
		Gender:  student.Gender,
		College: student.College,
		Major:   student.Major,
	}

	resp.StudentResp = studentResp

	selections, err := dao.StudentGetCourses(student.ID)
	if err != nil {
		return
	}

	var courseResps []model.CourseResp
	for _, selection := range selections {
		courseResp, err := GetCourseResp(selection.CourseID)
		if err != nil {
			return
		}
		courseResps = append(courseResps, courseResp)
	}

	return
}

func StudentDeleteCourse(c *gin.Context) (err error) {
	var selection model.CourseSelection
	uid, exists := c.Get("uid")
	if !exists {
		return errors.New("uid do not exists")
	}

	student, err := dao.GetStudentByUid(uid.(uint))
	selection.StudentID = student.ID

	courseIDStr := c.PostForm("courseID")
	courseID, err := strconv.Atoi(courseIDStr)
	selection.CourseID = uint(courseID)

	err = dao.StudentDeleteCourse(selection)
	return err
}
