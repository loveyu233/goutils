package email

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/smtp"
	"regexp"
	"strings"
	"time"
)

// IsValidEmail 校验email格式
func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regexpObj, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("无效的正则表达式:", err)
		return false
	}
	return regexpObj.MatchString(email)
}

// Email 发送者信息
type Email struct {
	SmtpHost       string //配置SMTP服务器地址和端口
	SmtpPort       int32
	SenderEmail    string // 配置发件人信息
	SenderPassword string
}

// RecipientEmail 接受者信息
type RecipientEmail struct {
	Email   string // 配置收件人信息
	Subject string // 构建邮件内容
	Body    string
}

func NewEmail(host string, port int32, email string, password string) *Email {
	return &Email{
		SmtpHost:       host,
		SmtpPort:       port,
		SenderEmail:    email,
		SenderPassword: password,
	}
}

func NewQQEmail(email string, password string) *Email {
	return &Email{
		SmtpHost:       "smtp.qq.com",
		SmtpPort:       587,
		SenderEmail:    email,
		SenderPassword: password,
	}
}

func New163Email(email string, password string) *Email {
	return &Email{
		SmtpHost:       "smtp.163.com",
		SmtpPort:       465,
		SenderEmail:    email,
		SenderPassword: password,
	}
}

// SendEmailBody 发送邮件
func (e *Email) SendEmailBody(rm RecipientEmail) error {
	// 创建认证
	auth := smtp.PlainAuth("", e.SenderEmail, e.SenderPassword, e.SmtpHost)

	// 构建发件人头部
	header := make(map[string]string)
	header["From"] = fmt.Sprintf("%s <%s>", "发件人姓名", e.SenderEmail)
	header["To"] = rm.Email
	header["Subject"] = rm.Subject

	var emailBody strings.Builder
	for key, value := range header {
		emailBody.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	emailBody.WriteString("\r\n")
	emailBody.WriteString(rm.Body)

	// 将邮件内容格式化为字节数组
	message := []byte(emailBody.String())

	// 发送邮件
	err := smtp.SendMail(fmt.Sprintf("%s:%d", e.SmtpHost, e.SmtpPort), auth, e.SenderEmail, []string{rm.Email}, message)
	return err
}

// SendEmailBodyAndSaveRedis 发送邮件并保存到redis中，保存value为发送的body
func (e *Email) SendEmailBodyAndSaveRedis(rm RecipientEmail, redisClient *redis.Client, ttl time.Duration) error {
	if err := e.SendEmailBody(rm); err != nil {
		return err
	}
	return redisClient.Set(context.TODO(), rm.Email, rm.Body, ttl).Err()
}
