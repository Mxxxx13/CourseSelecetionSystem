// @Title : student
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/16 10:52

package dao

import "CourseSeletionSystem/model"

func CreateStudentInfo(student model.Student) (err error) {
	err = DB.Create(&student).Error
	if err != nil {
		return err
	}
	return
}

func UpdateStudentInfo(student model.Student) (err error) {
	var dbStudent model.Student
	err = DB.Where("user_id = ?", student.UserID).First(&dbStudent).Error
	if err != nil {
		return err
	}

	err = DB.Save(&student).Error
	if err != nil {
		return err
	}
	return
}
