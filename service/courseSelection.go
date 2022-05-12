// @Title : courseSelection
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/17 18:37

package service

import (
	"encoding/json"
	"errors"
	"strconv"

	"CourseSeletionSystem/dao"
	"CourseSeletionSystem/model"
	"CourseSeletionSystem/rabbitmq"
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

	if !IsFindCourse(uint(courseID)) {
		return errors.New("所选的课不存在")
	}

	if !IsMaxNum(uint(courseID)) {
		return errors.New("选课人数已满")
	}

	selection.CourseID = uint(courseID)

	byteSelction, err := json.Marshal(selection)
	if err != nil {
		return
	}

	rabbitMQ := rabbitmq.NewRabbitMQWork("course_selection")
	rabbitMQ.PublishWork(string(byteSelction))
	//err = dao.StudentSelectCourse(selection)
	return
}

func IsFindCourse(cid uint) bool {
	_, err := dao.GetCourse(cid)
	if err != nil {
		return false
	}
	return true
}

func IsMaxNum(cid uint) bool {
	course, err := dao.GetCourse(cid)
	if err != nil {
		return false
	}

	if course.StuNum+1 > course.MaxNum {
		return false
	}

	return true
}

func StudentGetCourse(c *gin.Context) (resp model.StudentCourseResp, err error) {
	uid, exists := c.Get("uid")
	if !exists {
		return resp, errors.New("uid do not exists")
	}

	student, err := dao.GetStudentByUid(uid.(uint))
	studentResp := model.StudentResp{
		ID:      student.ID,
		Name:    student.Name,
		Number:  student.Number,
		Gender:  student.Gender,
		College: student.College,
		Major:   student.Major,
	}

	resp.StudentResp = studentResp

	selections, err := dao.GetSelectionBySid(student.ID)
	if err != nil {
		return
	}

	var courseResps []model.CourseResp
	for _, selection := range selections {
		courseResp, err := GetCourseResp(selection.CourseID)
		if err != nil {
			return resp, err
		}
		courseResps = append(courseResps, courseResp)
	}
	resp.CourseResps = courseResps

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
