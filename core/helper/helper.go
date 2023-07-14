package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	"go_disk/core/define"
	"math/rand"
	"net/smtp"
	"time"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity, name string) (string, error) {
	//id
	// identity
	//name
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		fmt.Println("这里也有")
		return "", err
	}
	return tokenString, nil
}

func MailSendCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "cjy <2812485760@qq.com>"
	e.To = []string{mail}
	e.Subject = "验证码测试"
	e.HTML = []byte("验证码是：<h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "2812485760@qq.com", "ugencsnkcxyadcif", "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		return err
	}
	return nil
}

// 生成包含数字和大写字母的4位验证码
func GetCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i <= define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
