// @Title : courseSelection
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/17 18:37

package dao

import "CourseSeletionSystem/model"

func StudentSelectCourse(selection model.CourseSelection) (err error) {
	err = DB.Create(&selection).Error
	return
}

func GetSelectionBySid(sid uint) (selections []model.CourseSelection, err error) {
	err = DB.Where("student_id = ?", sid).Find(&selections).Error
	return
}

func StudentDeleteCourse(selection model.CourseSelection) (err error) {
	err = DB.Where("student_id = ? and course_id = ?", selection.StudentID, selection.CourseID).
		Delete(&selection).Error
	return
}

func GetSelectionByCid(cid uint) (selections []model.CourseSelection, err error) {
	err = DB.Where("course_id = ?", cid).Find(&selections).Error
	return
}
