// @Title : student
// @Description :学生数据库层
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
	err = DB.Model(&student).Where("user_id = ?", student.UserID).Updates(map[string]interface{}{
		"name":    student.Name,
		"number":  student.Number,
		"gender":  student.Gender,
		"college": student.College,
		"major":   student.Major,
	}).Error
	if err != nil {
		return err
	}
	return
}
