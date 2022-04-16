// @Title : user
// @Description :用户逻辑层
// @Author : MX
// @Update : 2022/4/15 22:07

package service

import (
	"errors"

	"CourseSeletionSystem/dao"
	"CourseSeletionSystem/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context, role string) (err error) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 将密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("加密失败")
	}
	password = string(hashPassword)

	user := model.User{
		Username: username,
		Password: password,
		Role:     role,
	}

	err = dao.Register(user)
	return
}

// Login 将输入的password加密后和数据库中的password进行比较
// 返回uid和error, uid用于生成token
func Login(c *gin.Context) (uid uint, err error) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	pass, err := dao.Login(username) // 获取password
	if err != nil {
		return 0, errors.New("用户不存在")
	}
	uid, err = dao.GetUid(username) // 获取id
	if err != nil {
		return 0, errors.New("用户不存在")
	}

	// 对password进行验证
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(password))
	if err != nil {
		return 0, errors.New("密码错误")
	}
	return
}

// AlterUser
func AlterUser(c *gin.Context) (err error) {
	username := c.PostForm("username")
	id, exists := c.Get("uid")
	if !exists {
		return errors.New("id not exist")
	}
	err = dao.AlterUser(username, id.(uint))
	return
}

// TODO: 添加修改密码的功能
