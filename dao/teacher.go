// @Title : teacher
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/17 14:22

package dao

import "CourseSeletionSystem/model"

func CreateTeacher(teacher model.Teacher) (err error) {
	err = DB.Create(&teacher).Error
	return
}

func UpdateTeacher(id uint, teacher model.Teacher) (err error) {
	err = DB.Model(&teacher).Where("id = ?", id).Updates(map[string]interface{}{
		"name":    teacher.Name,
		"number":  teacher.Number,
		"gender":  teacher.Gender,
		"college": teacher.College,
	}).Error
	return
}

func GetTeacher(id uint) (teacher model.Teacher, err error) {
	err = DB.Where("id = ?", id).First(&teacher).Error
	return
}

func DeleteTeacher(id uint) (err error) {
	err = DB.Delete(&model.Teacher{}, id).Error
	return
}
