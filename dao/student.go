// @Title : student
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/16 10:52

package dao

import "CourseSeletionSystem/model"

func UpdateStudent(student model.Student) (err error) {
	err = DB.Model(&student).Updates(student).Error
	if err != nil {
		return err
	}
	return
}
