// @Title : student
// @Description :学生模型定义
// @Author : MX
// @Update : 2022/4/14 22:52

package model

import "github.com/jinzhu/gorm"

type Student struct {
	*gorm.Model
	Name    string // 名字
	Number  string // 学号
	Gender  string // 性别
	College string // 学院
	Major   string // 专业

	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
