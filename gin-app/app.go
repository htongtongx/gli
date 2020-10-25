package app

import (
	middleware "black-hole.com/modules/gin-app/middleware"
	"github.com/gin-gonic/gin"
)

func Fail(c *gin.Context, errCode int, errMsg string) {
	h := gin.H{
		"outCome": gin.H{"errorNum": errCode, "errorMessage": errMsg},
	}
	c.JSON(200, h)
	return
}

func Succ(c *gin.Context, data interface{}) {
	h := gin.H{
		"outCome": gin.H{"errorNum": 0},
	}
	if data != nil {
		h["data"] = data
	}
	c.JSON(200, h)
	return
}

func GetString(c *gin.Context, key, defaultValue string) string {
	data, exist := c.Get("data")
	if !exist {
		return defaultValue
	}
	temp := data.(*map[string]interface{})
	if (*temp)[key] == nil {
		return defaultValue
	}
	return (*temp)[key].(string)
}

func GetInt(c *gin.Context, key string, defaultValue int) (result int) {
	data, exist := c.Get("data")
	if !exist {
		return defaultValue
	}
	temp := data.(*map[string]interface{})
	if (*temp)[key] == nil {
		return defaultValue
	}
	return (*temp)[key].(int)
}

func Succ3(c *gin.Context, key string, data interface{}) {
	h := gin.H{
		"outCome": gin.H{"errorNum": 0},
	}
	if data != nil {
		h[key] = data
	}
	c.JSON(200, h)
	return
}

func AddPost(router gin.IRouter, group, url string, postData interface{}, f gin.HandlerFunc, noToken ...bool) {
	router.POST(url, f)
	if postData != nil {
		middleware.BindMap[group+url] = postData
	}
	if len(noToken) > 0 {
		if noToken[0] {
			middleware.IngoreURL[group+url] = 0
		}
	}
}

func AddPut(router gin.IRouter, group, url string, postData interface{}, f gin.HandlerFunc, noToken ...bool) {
	router.PUT(url, f)
	if postData != nil {
		middleware.BindMap[group+url] = postData
	}
	if len(noToken) > 0 {
		if noToken[0] {
			middleware.IngoreURL[group+url] = 0
		}
	}
}

func AddGet(router gin.IRouter, group, url string, f gin.HandlerFunc, noToken ...bool) {
	router.GET(url, f)
	if len(noToken) > 0 {
		if noToken[0] {
			middleware.IngoreURL[group+url] = 0
		}
	}
}

func AddDel(router gin.IRouter, group, url string, f gin.HandlerFunc, noToken ...bool) {
	router.DELETE(url, f)
	if len(noToken) > 0 {
		if noToken[0] {
			middleware.IngoreURL[group+url] = 0
		}
	}
}
