package sms

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
// 	"github.com/htongtongx/gli/redis"
// 	"github.com/htongtongx/gli/timer"
// 	"github.com/htongtongx/gli/util"
// )

// // type NasCarrierType int
// const (
// 	CacheToLocal = "local"
// 	CacheToRedis = "redis"
// )

// type AliSMSConf struct {
// 	RegionID     string `ini:"regionID"`     //区域ID
// 	AccessKeyID  string `ini:"accessKeyID"`  //访问权限ID
// 	AccessSecret string `ini:"accessSecret"` //访问权限秘钥
// 	SignName     string `ini:"signName"`     //短信模板签名
// 	TemplateCode string `ini:"templateCode"` //短信模板ID
// 	TemplateKey  string `ini:"templateKey"`  //短信模板自定义key
// 	CodeCount    int    `ini:"codeCount"`    //验证码位数
// 	Timeout      int    `ini:"timeout"`      //验证码过期时间 单位秒
// 	CacheTo      string `ini:"cacheTo"`      //验证码存储的位置“local”,"redis"
// }

// type AliSMS struct {
// 	cli   *dysmsapi.Client
// 	cfg   *AliSMSConf
// 	cache map[string]int
// 	redis *redis.Redis
// }

// func NewAliSMS(cfg *AliSMSConf, r *redis.Redis) (as *AliSMS) {
// 	fmt.Println(cfg)
// 	var err error
// 	as = &AliSMS{}
// 	as.cfg = cfg
// 	as.cli, err = dysmsapi.NewClientWithAccessKey(cfg.RegionID, cfg.AccessKeyID, cfg.AccessSecret)
// 	if err != nil {
// 		log.Println("初始化阿里云短信出错" + err.Error())
// 		return
// 	}
// 	if as.isCacheToLocal() {
// 		as.cache = make(map[string]int)
// 		fmt.Println("初始化map成功")
// 		go as.clear()
// 		return
// 	}
// 	if r == nil {
// 		log.Println("阿里云短信CacheTo=redis,redis实例不能为空")
// 	}
// 	as.redis = r
// 	return
// }

// func (a *AliSMS) isCacheToLocal() bool {
// 	return a.cfg.CacheTo == "" || a.cfg.CacheTo == CacheToLocal || a.cfg.CacheTo != CacheToRedis
// }

// func (a *AliSMS) clear() {
// 	t := time.NewTicker(1 * time.Hour)
// 	for {
// 		select {
// 		case <-t.C:
// 			now := timer.GetTimestamp()
// 			for k, v := range a.cache {
// 				if now > v {
// 					delete(a.cache, k)
// 				}
// 			}
// 		}
// 	}
// }

// // Send 发送验证码
// //错误码文档 https://help.aliyun.com/document_detail/101346.html?spm=a2c1g.8271268.10000.143.5d6edf25vaJLqE
// func (a *AliSMS) Send(phone string) (bool, error) {
// 	smsCode := util.RandNumByCount(a.cfg.CodeCount)
// 	request := dysmsapi.CreateSendSmsRequest()
// 	request.Scheme = "https"

// 	request.PhoneNumbers = phone
// 	request.SignName = a.cfg.SignName
// 	request.TemplateCode = a.cfg.TemplateCode
// 	request.TemplateParam = fmt.Sprintf(`{"%s":"%s"}`, a.cfg.TemplateKey, smsCode)
// 	// request.Headers["Content-Type"] = "application/json; charset=utf-8"

// 	response, err := a.cli.SendSms(request)
// 	if err != nil {
// 		return false, err
// 	}

// 	if response.Code != "OK" {
// 		return false, errors.New(response.Message)
// 	}
// 	err = a.saveTo(phone, smsCode)
// 	return true, err
// }

// func (a *AliSMS) saveTo(phone, smsCode string) (err error) {
// 	key := a.cacheKey(phone, smsCode)
// 	if a.isCacheToLocal() {
// 		a.cache[key] = timer.GetTimestamp() + a.cfg.Timeout
// 	} else {
// 		err = a.redis.SetEX(key, a.cfg.Timeout, "")
// 	}
// 	return
// }

// // Verify 当验证码不存在或超时时间超过当前时间戳，验证码无效
// func (a *AliSMS) Verify(phone, smsCode string) (bool, error) {
// 	key := a.cacheKey(phone, smsCode)
// 	if a.isCacheToLocal() {
// 		expires, ok := a.cache[key]
// 		return ok && timer.GetTimestamp() <= expires, nil
// 	}
// 	return a.redis.EXISTS(key)
// }

// func (a *AliSMS) cacheKey(phone, smsCode string) string {
// 	return fmt.Sprintf(`%s:%s`, phone, smsCode)
// }
