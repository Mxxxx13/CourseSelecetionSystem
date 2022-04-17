// @Title : course
// @Description :课程逻辑层
// @Author : MX
// @Update : 2022/4/16 20:58

package service

import (
	"errors"
	"log"
	"strconv"

	"CourseSeletionSystem/dao"
	"CourseSeletionSystem/model"
	"github.com/gin-gonic/gin"
)

func CreateCourse(c *gin.Context) (err error) {
	var course model.Course

	if err = c.ShouldBind(&course); err != nil {
		return err
	}

	err = dao.CreateCourse(course)
	return
}

func UpdateCourse(c *gin.Context) (err error) {
	var course model.Course

	if err = c.ShouldBind(&course); err != nil {
		return err
	}

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	err = dao.UpdateCourse(uint(id), course)
	return err
}

func GetCourse(c *gin.Context) (courseResp model.CourseResp, err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return courseResp, errors.New("请求参数错误")
	}
	// 查询用户
	course, err := dao.GetCourse(uint(id))
	teacher, err := dao.GetTeacher(course.TeacherID)
	courseResp = model.CourseResp{
		Name:    course.Name,
		Score:   course.Score,
		MaxNum:  course.MaxNum,
		StuNum:  course.StuNum,
		Teacher: teacher,
	}
	return
}

func DeleteCourse(c *gin.Context) (err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	err = dao.DeleteCourse(uint(id))
	return err
}
