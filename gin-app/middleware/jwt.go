package middleware

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/htongtongx/gli/gconf"
)

var IngoreURL = make(map[string]int)
var tokenHeader = "Authorization"

func GetToken(c *gin.Context, header string) string {
	if header == "" {
		header = tokenHeader
	}
	token := c.GetHeader(header)
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.Replace(token, "Bearer ", "", 1)
	}
	return token
}

func AuthToken(jwtCfg *gconf.JWTConf) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !jwtCfg.Enabled {
			return
		}
		_, ok := IngoreURL[c.FullPath()]
		if !ok {
			authKey := GetToken(c, jwtCfg.Header)
			t, err := ParseToken(authKey, jwtCfg.Secrets)
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
