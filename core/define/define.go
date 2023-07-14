package define

import "github.com/dgrijalva/jwt-go"

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "go_disk_key"

// 定义验证码的长度
var CodeLength = 6

// 定义验证码的过期时间
var CodeExpire = 300
