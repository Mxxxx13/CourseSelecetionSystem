// @Title : teacher
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/17 14:37

package service

import (
	"errors"
	"log"
	"strconv"

	"CourseSeletionSystem/dao"
	"CourseSeletionSystem/model"
	"github.com/gin-gonic/gin"
)

func CreateTeacher(c *gin.Context) (err error) {
	var teacher model.Teacher

	if err = c.ShouldBind(&teacher); err != nil {
		return err
	}

	uid, exsits := c.Get("uid")
	if !exsits {
		return errors.New("uid not exists")
	}
	teacher.UserID = uid.(uint)

	err = dao.CreateTeacher(teacher)
	return
}

func UpdateTeacher(c *gin.Context) (err error) {
	var teacher model.Teacher

	if err = c.ShouldBind(&teacher); err != nil {
		return err
	}

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	err = dao.UpdateTeacher(uint(id), teacher)
	return err
}

func GetTeacher(c *gin.Context) (teacherResp model.TeacherResp, err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return teacherResp, errors.New("请求参数错误")
	}

	teacherResp, err = GetTeacherResp(uint(id))
	return
}

func GetTeacherResp(id uint) (teacherResp model.TeacherResp, err error) {
	student, err := dao.GetTeacher(id)
	teacherResp = model.TeacherResp{
		Name:    student.Name,
		Number:  student.Number,
		Gender:  student.Gender,
		College: student.College,
	}
	return
}

func DeleteTeacher(c *gin.Context) (err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	err = dao.DeleteTeacher(uint(id))
	return err
}

func GetStudentSelection(c *gin.Context) (resp model.TeacherStudentSelectionResp, err error) {
	uid, exsits := c.Get("uid")
	if !exsits {
		return resp, errors.New("uid not exists")
	}

	teacher, err := dao.GetTeacherByUid(uid.(uint))
	if err != nil {
		return
	}

	resp.TeacherResp, err = GetTeacherResp(teacher.UserID)
	if err != nil {
		return
	}

	courses, err := dao.GetCourseByTid(teacher.ID)
	if err != nil {
		return
	}

	for _, course := range courses {
		var ssResp model.StudentSelectionResp
		courseResp, err := GetCourseResp(course.ID)
		if err != nil {
			return resp, err
		}
		ssResp.CourseResp = courseResp

		selections, err := dao.GetSelectionByCid(course.ID)
		for _, selection := range selections {
			studentResp, err := GetStudentResp(selection.StudentID)
			if err != nil {
				return resp, err
			}

			ssResp.StudentsResp = append(ssResp.StudentsResp, studentResp)
		}
		resp.StudentSelectionResp = append(resp.StudentSelectionResp, ssResp)
	}
	return
}
