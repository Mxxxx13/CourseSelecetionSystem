// @Title : user
// @Description :用户数据库层
// @Author : MX
// @Update : 2022/4/15 22:06

package dao

import "CourseSeletionSystem/model"

//Register 将username,password插入数据库
func Register(user model.User) (err error) {
	if err = DB.Create(&user).Error; err != nil {
		return
	}
	return
}

// GetUser 获取user对象
func GetUser(username string) (user model.User, err error) {
	if err = DB.Where("username = ?", username).First(&user).Error; err != nil {
		return
	}
	return
}

// AlterUser 更新用户信息
func AlterUser(username string, uid uint) (err error) {

	if err = DB.Model(&model.User{}).Where("id = ?", uid).Update("username", username).Error; err != nil {
		return
	}
	return
}

func AlterUserPassword(password string, uid uint) (err error) {
	if err = DB.Model(&model.User{}).Where("id = ?", uid).Update("password", password).Error; err != nil {
		return
	}
	return
}

func GetUserByEmail(address string) (user model.User, err error) {
	if err = DB.Model(&user).Where("email = ?", address).First(&user).Error; err != nil {
		return
	}
	return
}
