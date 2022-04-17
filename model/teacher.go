// @Title : teacher
// @Description :教师模型定义
// @Author : MX
// @Update : 2022/4/14 23:17

package model

import "github.com/jinzhu/gorm"

type Teacher struct {
	*gorm.Model
	Name    string // 名字
	Number  string
	Gender  string // 性别
	College string // 学院

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type TeacherResp struct {
	Name    string // 名字
	Number  string
	Gender  string // 性别
	College string // 学院
}
