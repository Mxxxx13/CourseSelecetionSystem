// @Title : user
// @Description :用户模型定义
// @Author : MX
// @Update : 2022/4/14 22:23

package model

import "github.com/jinzhu/gorm"

type User struct {
	*gorm.Model
	Username string `gorm:"unique"` // 用户名
	Password string // 密码
	//Email    string
	Role string // 身份
}
