// @Title : student
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/16 10:52

package service

import (
	"errors"
	"log"
	"strconv"

	"CourseSeletionSystem/dao"
	"CourseSeletionSystem/model"
	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) (err error) {
	var student model.Student

	if err = c.ShouldBind(&student); err != nil {
		return err
	}

	uid, exsits := c.Get("uid")
	if !exsits {
		return errors.New("uid not exists")
	}
	student.UserID = uid.(uint)

	err = dao.CreateStudent(student)
	return
}

func UpdateStudent(c *gin.Context) (err error) {
	var student model.Student

	if err = c.ShouldBind(&student); err != nil {
		return err
	}

	uid, exsits := c.Get("uid")
	if !exsits {
		return errors.New("uid not exists")
	}
	student.UserID = uid.(uint)

	err = dao.UpdateStudent(student)
	return err
}

func GetStudent(c *gin.Context) (studentResp model.StudentResp, err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return studentResp, errors.New("请求参数错误")
	}
	// 查询用户
	student, err := dao.GetStudent(uint(id))
	studentResp = model.StudentResp{
		Name:    student.Name,
		Number:  student.Number,
		Gender:  student.Gender,
		College: student.College,
		Major:   student.Major,
	}
	return
}

func DeleteStudent(c *gin.Context) (err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	err = dao.DeleteStudent(uint(id))
	return err
}
