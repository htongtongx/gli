package sms

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"io/ioutil"
// 	"time"

// 	"github.com/astaxie/beego/httplib"
// 	"github.com/htongtongx/gli/cryptos"
// )

// const apiurl = "http://www.zhiliandaili.com/Wap/App/init"
// const pwd = "dUkks43k"
// const smstype = "wanchen_code"

// const urlkey = "url_list"
// const sendret = "sendsms_zhilian"
// const checksum = "checksms"

// var (
// 	urls = make(map[string]interface{})
// )

// type misi struct {
// 	Phone    string `phone`
// 	Type     string `type`
// 	Send_pwd string `send_pwd`
// 	Limit    string `limit`
// 	Token    string `token`
// }
// type checkrep struct {
// 	SmsNO string
// 	Code  string
// }

// func init() {
// 	res, err := send("", misi{})
// 	if err == nil {

// 		err = json.Unmarshal(res, &urls)

// 		if err != nil {

// 			fmt.Println(err)
// 		}
// 	}
// }
// func send(url string, data interface{}) (ret []byte, err error) {
// 	if url == "" {
// 		resp, er := httplib.Get(apiurl).SetTimeout(10*time.Second, 5*time.Second).Response()
// 		if er != nil {
// 			return nil, er
// 		}
// 		return ioutil.ReadAll(resp.Body)
// 	}
// 	req := httplib.Post(url)

// 	switch data.(type) {
// 	case misi:
// 		req.Param("phone", data.(misi).Phone)
// 		req.Param("type", data.(misi).Type)
// 		req.Param("send_pwd", data.(misi).Send_pwd)
// 		req.Param("limit", data.(misi).Limit)
// 		break
// 	case checkrep:
// 		v1 := data.(checkrep)
// 		req.Param("sms_no", v1.SmsNO)
// 		req.Param("code", v1.Code)
// 		break
// 	}

// 	//req.SetTimeout(10*time.Second,5*time.Second).Header("XX-Token",v.Token)
// 	req.SetTimeout(10*time.Second, 5*time.Second)
// 	rep, er := req.Response()
// 	if er != nil {
// 		return nil, er
// 	}
// 	if rep.StatusCode != 200 {
// 		return nil, errors.New(fmt.Sprintf("error:[%d]", rep.StatusCode))
// 	}
// 	return ioutil.ReadAll(rep.Body)
// }

// func geturl(skey string) (ret string, err error) {

// 	urllist := urls[urlkey].(map[string]interface{})

// 	for key, url := range urllist {

// 		//fmt.Println(fmt.Sprintf("key:%s,url:%s",key,url))

// 		if key == skey {
// 			return url.(string), nil
// 		}
// 	}
// 	return "", errors.New("get url error")
// }

// var smsMap = make(map[string]string)

// func Send(phone string) (err error) {
// 	token := cryptos.MD5(fmt.Sprintf("%v", time.Now().Unix()))
// 	data := misi{
// 		phone,
// 		smstype,
// 		pwd,
// 		"1",
// 		token,
// 	}

// 	url, err := geturl(sendret)
// 	if err != nil {
// 		return
// 	}
// 	var ret []byte
// 	ret, err = send(url, data)
// 	if err != nil {
// 		return
// 	}
// 	rr := make(map[string]interface{})
// 	if err = json.Unmarshal(ret, &rr); err != nil {
// 		return
// 	}
// 	if rr["code"].(float64) == 0 {
// 		smsMap[phone] = rr["data"].(string)
// 	}
// 	return
// }

// //parmes code 手机验证码
// //parmes phone  手机号
// //parmes SmsNO  是发送手机验证码成功后返回的data
// func Check(code string, phone string) (resultMap map[string]interface{}, err error) {
// 	token, ok := smsMap[phone]
// 	if !ok {
// 		return nil, errors.New("手机验证码失败")
// 	}
// 	v, er := geturl(checksum)
// 	if er != nil {
// 		return nil, er
// 	}
// 	var result []byte
// 	result, err = send(v, checkrep{SmsNO: token, Code: code})
// 	resultMap = make(map[string]interface{})
// 	if err = json.Unmarshal(result, &resultMap); err != nil {
// 		return
// 	}
// 	return resultMap, err
// }
