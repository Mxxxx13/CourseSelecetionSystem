// @Title : user
// @Description :用户逻辑层
// @Author : MX
// @Update : 2022/4/15 22:07

package service

import (
	"errors"
	"strconv"

	"CourseSeletionSystem/dao"
	"CourseSeletionSystem/model"
	"CourseSeletionSystem/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context, role string) (err error) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 将密码加密
	password, err = util.EncryptionPassword(password)
	if err != nil {
		return
	}

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
func Login(c *gin.Context) (user model.User, err error) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err = dao.GetUser(username)
	if err != nil {
		return user, errors.New("用户不存在")
	}

	// 对password进行验证
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("密码错误")
	}
	return
}

func AlterUser(c *gin.Context) (err error) {
	username := c.PostForm("username")
	id, exists := c.Get("uid")
	if !exists {
		return errors.New("id not exist")
	}
	err = dao.AlterUser(username, id.(uint))
	return
}

func CheckCode(c *gin.Context) (err error) {
	uid, _ := c.Get("uid")
	key := "code:" + uid.(string)
	val, err := dao.Redis.Get(key).Result()
	code := c.PostForm("code")
	if code != val {
		return errors.New("验证码错误")
	}
	return
}

func AlterPassword(c *gin.Context) (err error) {
	err = CheckCode(c)
	if err != nil {
		return
	}

	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirmPassword")

	if password != confirmPassword {
		return errors.New("两次输入的密码不一样")
	}

	// 将密码加密
	password, err = util.EncryptionPassword(password)
	if err != nil {
		return
	}

	id, exists := c.Get("uid")
	if !exists {
		return errors.New("id not exist")
	}

	err = dao.AlterUserPassword(password, id.(uint))
	return
}

func SendEmail(c *gin.Context) (err error) {
	address := c.PostForm("address") // 邮箱地址

	code := util.GenerateCode()

	user, err := dao.GetUserByEmail(address)
	if err != nil {
		return errors.New("输入的邮箱错误")
	}
	key := "code:" + strconv.Itoa(int(user.ID))
	// 将验证码存入redis
	// 有效期5分钟
	err = dao.Redis.Set(key, code, 300).Err()
	if err != nil {
		return err
	}

	c.Set("uid", user.ID)

	// 通过邮件发送验证码
	err = util.SendEmail(code, address)
	if err != nil {
		return errors.New("发送邮件失败")
	}
	return nil
}
