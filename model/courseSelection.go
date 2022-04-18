// @Title : courseSeletion
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/4/17 16:59

package model

import "github.com/jinzhu/gorm"

type CourseSelection struct {
	*gorm.Model

	StudentID uint
	Student   Student

	CourseID uint
	Course   Course
}
