// @Title : student
// @Description :学生逻辑层
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

	err = StudentCheck(c)
	if err != nil {
		return
	}

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	err = dao.UpdateStudent(uint(id), student)
	return err
}

func GetStudent(c *gin.Context) (studentResp model.StudentResp, err error) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return studentResp, errors.New("请求参数错误")
	}
	studentResp, err = GetStudentResp(uint(id))
	return
}

func GetStudentResp(sid uint) (studentResp model.StudentResp, err error) {
	student, err := dao.GetStudent(sid)
	studentResp = model.StudentResp{
		ID:      student.ID,
		Name:    student.Name,
		Number:  student.Number,
		Gender:  student.Gender,
		College: student.College,
		Major:   student.Major,
	}
	return
}

func DeleteStudent(c *gin.Context) (err error) {
	err = StudentCheck(c)
	if err != nil {
		return
	}

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	err = dao.DeleteStudent(uint(id))
	return err
}

// StudentCheck 对用户身份进行校验
func StudentCheck(c *gin.Context) (err error) {
	role, exists := c.Get("role")
	if !exists {
		return errors.New("role not exists")
	}
	if role.(string) == "admin" {
		return nil
	}

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return
	}

	uid, exists := c.Get("uid")
	if !exists {
		return errors.New("uid not exists")
	}

	student, err := dao.GetStudentByUid(uid.(uint))

	if uint(id) != student.ID {
		return errors.New("权限不够")
	}
	return
}
