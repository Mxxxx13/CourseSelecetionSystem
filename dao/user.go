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

// Login 查询password并返回
func Login(username string) (password string, err error) {
	var user model.User
	if err = DB.Where("username = ?", username).First(&user).Error; err != nil {
		return
	}
	return user.Password, nil
}

// GetUid 根据username查询uid
func GetUid(username string) (uid uint, err error) {
	var user model.User
	if err = DB.Where("username = ?", username).First(&user).Error; err != nil {
		return
	}
	return user.ID, nil
}

// GetUser 获取user对象
func GetUser(uid uint) (user model.User, err error) {
	if err = DB.Where("id = ?", uid).First(&user).Error; err != nil {
		return
	}
	return
}

// AlterUser 更新用户信息
func AlterUser(username string, uid uint) (err error) {
	var user model.User
	if err = DB.Model(&user).Where("id = ?", uid).Update("username", username).Error; err != nil {
		return
	}
	return
}
