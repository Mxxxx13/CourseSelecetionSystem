// @Title : student
// @Description :学生模型定义
// @Author : MX
// @Update : 2022/4/14 22:52

package model

import "github.com/jinzhu/gorm"

type Student struct {
	*gorm.Model
	Name    string `form:"name"`    // 名字
	Number  string `form:"number"`  // 学号
	Gender  string `form:"gender"`  // 性别
	College string `form:"college"` // 学院
	Major   string `form:"major"`   // 专业

	UserID uint `gorm:"unique"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type StudentResp struct {
	Name    string `form:"name"`    // 名字
	Number  string `form:"number"`  // 学号
	Gender  string `form:"gender"`  // 性别
	College string `form:"college"` // 学院
	Major   string `form:"major"`   // 专业
}

type StudentCourseResp struct {
	StudentResp StudentResp
	CourseResps []CourseResp
}
