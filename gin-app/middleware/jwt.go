package middleware

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var IngoreURL = make(map[string]int)
var TokenHeader = "Authorization"

func AuthToken(secrets, headerName string, isDev bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ignore := isIgnore(c.FullPath())
		_, ok := IngoreURL[c.FullPath()]
		if !ok {
			authKey := c.GetHeader(headerName)
			if strings.HasPrefix(authKey, "Bearer ") {
				authKey = strings.Replace(authKey, "Bearer ", "", 1)
			}
			t, err := ParseToken(authKey, secrets)
			if err != nil {
				c.AbortWithStatusJSON(200, gin.H{"code": 101, "msg": "token invalid" + err.Error()})
				return
			}
			claim := map[string]interface{}(t.Claims.(jwt.MapClaims))
			c.Set("UserID", int(claim["user_id"].(float64)))
			c.Set("UserName", claim["user_name"].(string))
			c.Set("tokenMap", claim)
		}
	}
}

func GenAuthToken(ueseID int, userName string, secrets string, exp int64, ext ...map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	// b := make(map[string]interface{})
	// claims = jwt.MapClaims(b)
	//添加令牌期限
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(exp)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["user_id"] = ueseID
	claims["user_name"] = userName
	if len(ext) > 0 {
		for k, v := range ext[0] {
			claims[k] = v
		}
	}
	token.Claims = claims
	return token.SignedString([]byte(secrets))
}

func ParseToken(token string, secrets string) (*jwt.Token, error) {
	return jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(secrets), nil
	})
}
