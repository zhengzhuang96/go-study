// Author: zhengzhuang
// Date: 2021-07-27 15:48:29
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-27 15:49:02
// Description: In User Settings Edit
// FilePath: /go-study/gin/gin-jwt-demo/tool/JWTClaims.go
package tool

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

func GetSignKey() string {
	return SignKey
}

type JWT struct {
	SigningKey []byte
}

type JwtClaims struct {
	ID    string `json:"userId"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	// jwt-go 提供的标准claim
	jwt.StandardClaims
}

var (
	// 自定义的token密钥
	secret = []byte("16849841325189456f487")
	// // 该路由下不效验token
	// noVerify = []interface{}{"/login", "/ping"}
	// // token 有效时间(纳秒)
	// effectTime              = 2 * time.Hour
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "newtrekWang"
)

// 生成token
func (j *JWT) CreateToken(jwtC JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwtC)
	return token.SignedString(j.SigningKey)
}

// 效验token
// func (j *JWT) JwtVerify(c *gin.Context) (*JwtClaims, error) {
// 	//过滤是否验证token
// 	// if utils.IsContainArr(noVerify, c.Request.RequestURI) {
// 	// 	return
// 	// }
// 	return j.parseToken(token)
// }

// 解析/效验Token
func (j *JWT) ParseToken(tokenString string) (*JwtClaims, error) {
	// token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	return j.SigningKey, nil
	// })
	// if err != nil {
	// 	if ve, ok := err.(*jwt.ValidationError); ok {
	// 		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
	// 			return nil, TokenMalformed
	// 		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
	// 			// Token is expired
	// 			return nil, TokenExpired
	// 		} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
	// 			return nil, TokenNotValidYet
	// 		} else {
	// 			return nil, TokenInvalid
	// 		}
	// 	}
	// }
	// if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
	// 	return claims, nil
	// }
	// return nil, TokenInvalid

	//解析token
	// println(c.Request.RequestURI)
	// println(c.Request.RequestURI)
	// if utils.IsContainArr(noVerify, c.Request.RequestURI) {
	// 	return
	// }
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		panic("token is valid")
	}
	return claims, nil
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
