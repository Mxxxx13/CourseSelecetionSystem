// @Title : sendEmail
// @Description :发送邮件
// @Author : MX
// @Update : 2022/4/21 20:17

package util

import (
	"fmt"
	"log"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

var pool *email.Pool

func NewEmailPool() {
	p, err := email.NewPool(
		"smtp.qq.com:25",
		4,
		smtp.PlainAuth("", "1623422215@qq.com", "nttjeuimfnzrdgda", "smtp.qq.com"),
	)

	if err != nil {
		log.Fatal("failed to create pool:", err)
	}
	pool = p
	return
}

// SendEmail 发送邮件
func SendEmail(code string, address string) error {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "选课系统<1623422215@qq.com>"
	// 设置接收方的邮箱
	e.To = []string{address}
	//设置主题
	e.Subject = "验证码"

	html := fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>你本次的验证码为%s,为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>`, code)

	//设置文件发送的内容
	e.HTML = []byte(html)

	err := pool.Send(e, 10*time.Second)
	if err != nil {
		return err
	}
	return nil
}
