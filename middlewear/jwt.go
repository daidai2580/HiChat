package middlewear

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var (
	TokenExpired = errors.New("Token is expired.")
)

// 指定秘钥
var jwtSecret = []byte("jhome")

type Claims struct {
	UserID uint `json:"userId"`
	jwt.StandardClaims
}

func GenerateToken(userId uint, iss string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(48 * 30 * time.Hour)
	claim := Claims{UserID: userId, StandardClaims: jwt.StandardClaims{ExpiresAt: expireTime.Unix(), Issuer: iss}}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tokenClaim.SignedString(jwtSecret)
	return token, err
}

func parseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func JWY() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")
		user := c.PostForm("id")
		userId, err := strconv.Atoi(user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "您userId不合法",
			})
			c.Abort()
			return
		}
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "请登录",
			})
			c.Abort()
			return
		} else {
			claims, err := parseToken(token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "token失效",
				})
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				err = TokenExpired
				c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "授权已过期",
				})
				c.Abort()
				return
			}
			if claims.UserID != uint(userId) {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "您的登录不合法",
				})
				c.Abort()
				return
			}
			fmt.Print("token 认证成功")
			c.Next()
		}
	}
}
