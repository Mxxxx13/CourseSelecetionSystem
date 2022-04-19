// @Title : course
// @Description :课程模型定义
// @Author : MX
// @Update : 2022/4/15 20:34

package model

import (
	"github.com/jinzhu/gorm"
)

type Course struct {
	*gorm.Model
	Name   string  `form:"name"`   // 名字
	Score  float32 `form:"score"`  // 学分
	MaxNum uint    `form:"maxnum"` // 课程人数上限
	StuNum uint    `form:"stunum"` // 选课学生数量
	Time   string  `form:"time"`   // 上课时间
	Week   string  `form:"week"`   // 上课周

	TeacherID uint `form:"teacherID"` // 任课老师id
}

type CourseResp struct {
	Name        string
	Score       float32
	MaxNum      uint
	StuNum      uint
	Time        string
	Week        string
	TeacherResp TeacherResp
}
