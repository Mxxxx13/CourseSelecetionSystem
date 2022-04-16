// @Title : student
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/16 10:52

package service

import (
	"errors"

	"CourseSeletionSystem/dao"
	"CourseSeletionSystem/model"
	"github.com/gin-gonic/gin"
)

func StudentInfo(c *gin.Context) (err error) {
	var student model.Student
	if err = c.ShouldBind(&student); err != nil {
		return err
	}

	uid, exsits := c.Get("uid")
	if !exsits {
		return errors.New("uid not exists")
	}

	student.User.ID = uid.(uint)

	err = dao.UpdateStudent(student)
	return err
}
