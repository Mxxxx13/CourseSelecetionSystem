// @Title : student
// @Description :学生数据库层
// @Author : MX
// @Update : 2022/4/16 10:52

package dao

import "CourseSeletionSystem/model"

func CreateStudent(student model.Student) (err error) {
	err = DB.Create(&student).Error
	return
}

func UpdateStudent(id uint, student model.Student) (err error) {
	err = DB.Model(&student).Where("id = ?", id).Updates(map[string]interface{}{
		"name":    student.Name,
		"number":  student.Number,
		"gender":  student.Gender,
		"college": student.College,
		"major":   student.Major,
	}).Error
	return
}

func GetStudent(id uint) (student model.Student, err error) {
	err = DB.Where("id = ?", id).First(&student).Error
	return
}

func GetStudentByUid(uid uint) (student model.Student, err error) {
	err = DB.Where("user_id = ?", uid).First(&student).Error
	return
}

func DeleteStudent(id uint) (err error) {
	err = DB.Unscoped().Delete(&model.Student{}, id).Error
	return
}
