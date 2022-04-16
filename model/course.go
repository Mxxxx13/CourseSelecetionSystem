// @Title : course
// @Description :课程模型定义
// @Author : MX
// @Update : 2022/4/15 20:34

package model

import "github.com/jinzhu/gorm"

type Course struct {
	*gorm.Model
	Name   string // 名字
	Score  int    // 学分
	MaxNum uint   // 课程人数上限
	StuNum uint   // 选课学生数量
}
