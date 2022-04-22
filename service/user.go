// @Title : user
// @Description :用户逻辑层
// @Author : MX
// @Update : 2022/4/15 22:07

package service

import (
	"errors"
	"strconv"
	"time"

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

func BindEmail(c *gin.Context) (err error) {
	address := c.PostForm("address")
	id := c.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("请求参数错误")
	}

	err = CheckCode(c)
	if err != nil {
		return
	}

	err = dao.BindEmail(address, uint(uid))
	return
}

func CheckCode(c *gin.Context) (err error) {
	uid := c.Param("id")
	key := "code:" + uid
	val, err := dao.Redis.Get(key).Result()
	code := c.PostForm("code")
	if code != val {
		return errors.New("验证码错误")
	}
	return
}

func AlterPassword(c *gin.Context) (err error) {
	id := c.Param("id")
	userid, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("请求参数错误")
	}
	uid := uint(userid)

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

	err = dao.AlterUserPassword(password, uid)
	return
}

func SendEmail(c *gin.Context) (err error) {
	address := c.PostForm("address") // 邮箱地址

	id := c.Param("id")
	userid, err := strconv.Atoi(id)
	uid := uint(userid)
	if err != nil {
		return errors.New("请求参数错误")
	}

	err = SendEmailToUser(address, uid)
	if err != nil {
		return
	}
	return
}

func SendEmailToUser(address string, uid uint) (err error) {
	code := util.GenerateCode()

	// 将验证码存入redis
	// 有效期5分钟
	key := "code:" + strconv.Itoa(int(uid))
	err = dao.Redis.Set(key, code, 300*time.Second).Err()
	if err != nil {
		return err
	}

	err = util.SendEmail(code, address)
	if err != nil {
		return errors.New("发送邮件失败")
	}
	return
}
