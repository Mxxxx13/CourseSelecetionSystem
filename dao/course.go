// @Title : course
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/16 20:58

package dao

import (
	"CourseSeletionSystem/model"
)

func CreateCourse(course model.Course) (err error) {
	err = DB.Create(&course).Error
	return
}

func UpdateCourse(id uint, course model.Course) (err error) {
	err = DB.Model(&course).Where("id = ?", id).Updates(map[string]interface{}{
		"name":       course.Name,
		"score":      course.Score,
		"maxnum":     course.MaxNum,
		"teacher_ID": course.TeacherID,
		"time":       course.Time,
		"week":       course.Week,
	}).Error
	return
}

func GetCourse(id uint) (course model.Course, err error) {
	err = DB.Where("id = ?", id).First(&course).Error
	return
}

func DeleteCourse(id uint) (err error) {
	err = DB.Delete(&model.Course{}, id).Error
	return
}
