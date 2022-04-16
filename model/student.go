// @Title : student
// @Description :学生模型定义
// @Author : MX
// @Update : 2022/4/14 22:52

package model

import "github.com/jinzhu/gorm"

type Student struct {
	*gorm.Model
	Name    string `json:"name"`    // 名字
	Number  string `json:"number"`  // 学号
	Gender  string `json:"gender"`  // 性别
	College string `json:"college"` // 学院
	Major   string `json:"major"`   // 专业

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
