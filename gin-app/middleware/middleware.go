package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var BindMap = make(map[string]interface{})

// 处理post的data中间件
func ParseData() gin.HandlerFunc {
	return func(c *gin.Context) {
		if (c.Request.Method == "POST" || c.Request.Method == "PUT") && c.ContentType() == string(binding.MIMEJSON) {
			var data interface{}
			postStruct, ok := BindMap[c.FullPath()]
			if !ok {
				return
			}
			data = reflect.New(reflect.TypeOf(postStruct)).Interface()
			err := c.ShouldBindJSON(data)
			if err != nil {
				c.AbortWithStatusJSON(200, gin.H{"code": 100, "msg": err.Error()})
				return
			}
			c.Set("data", data)
		}
	}
}

//打印body的中间件
func PrintBody() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := ctx.GetRawData()
		if err != nil {
			log.Println(err.Error())
		}

		// log.Printf("data: %v\n", ctx.FullPath())

		log.Printf("%s: %s\n", ctx.Request.Method, ctx.Request.URL.String())
		if len(data) > 0 {
			log.Printf("data: %v\n", string(data))
		}

		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 关键点
		ctx.Next()
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func PrintResp() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		ctx.Next()
		log.Println(blw.body.String())
	}
}

func getArgs(c *gin.Context) []byte {
	if c.ContentType() == "multipart/form-data" {
		c.Request.ParseMultipartForm(1024)
	} else {
		c.Request.ParseForm()
	}
	args, _ := json.Marshal(c.Request.Form)
	return args
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		// c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,Access-Control-Allow-Origin,access-control-allow-headers")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
