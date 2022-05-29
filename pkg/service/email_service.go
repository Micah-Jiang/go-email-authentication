package service

import (
	"crypto/tls"
	"fmt"
	"go-email-authentication/pkg/cache"
	"go-email-authentication/pkg/utils"
	"gopkg.in/gomail.v2"
	"math/rand"
	"net/http"
	"time"
)

func SendEmail(username, email string) (int, string) {
	err := SendMessage(username, email)
	if err != nil {
		return http.StatusInternalServerError, fmt.Sprintf("send email failed, errMsg: %s", err)
	}
	return http.StatusOK, "send email success"
}

func SendMessage(receiverName, receiverEmail string) error {
	//message to html
	message := `
    <p> Hey %s,</p>
	
		<p style="text-indent:2em"><h2>%s</h2>, the verify code of your account, it will be expired after 1 minute.</p> 

	<p style="text-indent:2em">Best wishes!</p>
	`

	// QQ 邮箱：
	// SMTP 服务器地址：smtp.qq.com（SSL协议端口：465/994 | 非SSL协议端口：25）
	// 163 邮箱：
	// SMTP 服务器地址：smtp.163.com（端口：25）
	host := utils.GetInterfaceToString(utils.Get("email.mail.host"))
	port := 25
	userName := utils.GetInterfaceToString(utils.Get("email.mail.username"))
	password := utils.GetInterfaceToString(utils.Get("email.mail.password"))
	from := utils.GetInterfaceToString(utils.Get("email.mail.from"))
	subject := utils.GetInterfaceToString(utils.Get("email.mail.subject"))

	m := gomail.NewMessage()
	m.SetHeader("From", from+"<"+userName+">")                            // 增加发件人别名
	m.SetHeader("To", receiverEmail)                                      // 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
	m.SetHeader("Subject", "["+subject+"] "+"please verify your account") // 邮件主题

	// generate verify code
	randCode := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	// store verify code in redis
	err := cache.SetSendPhoneCodeCache(receiverEmail, randCode)
	if err != nil {
		return err
	}

	// text/html 的意思是将文件的 content-type 设置为 text/html 的形式，浏览器在获取到这种文件时会自动调用html的解析器对文件进行相应的处理。
	// 可以通过 text/html 处理文本格式进行特殊处理，如换行、缩进、加粗等等
	m.SetBody("text/html", fmt.Sprintf(message, receiverName, randCode))

	// text/plain的意思是将文件设置为纯文本的形式，浏览器在获取到这种文件时并不会对其进行处理
	// m.SetBody("text/plain", "纯文本")
	// m.Attach("test.sh")   // 附件文件，可以是文件，照片，视频等等
	// m.Attach("lolcatVideo.mp4") // 视频
	// m.Attach("lolcat.jpg") // 照片

	d := gomail.NewDialer(
		host,
		port,
		userName,
		password,
	)

	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
