package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/golang-module/carbon"
	"myBlog/models"
	"myBlog/setting"
	"strconv"
)

var (
	jwtKey = setting.Config.JwtKey
)

//GetToken 通过用户ID获取token
func GetToken(uid int, email string) (token string, expiresTime string, err error) {
	now := carbon.Now()
	expiresAt := now.AddHours(6)
	claims := models.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Carbon2Time().Unix(),
			IssuedAt:  now.Carbon2Time().Unix(),
			Issuer:    "root",
			Id:        strconv.Itoa(uid),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString(jwtKey)
	expiresTime = expiresAt.Format("Y-m-d H:i")
	return
}

func ParseToken(tokenStr string) (*jwt.Token, *models.Claims, error) {
	claim := models.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, &claim, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, &claim, err
}
