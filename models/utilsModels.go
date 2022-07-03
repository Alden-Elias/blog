package models

import "github.com/dgrijalva/jwt-go"

//Claims 声明， 保存在token中的字段
type Claims struct {
	Email string
	jwt.StandardClaims
}
