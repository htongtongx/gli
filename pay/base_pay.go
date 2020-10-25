package pay

import (
	"fmt"
	"time"
)

const Pay_Alipay = 1
const Pay_Wechat = 2
const Pay_Alipay_Prix = "alipay"
const Pay_Wechat_Prix = "wechat"

type PayClientApi struct {
	AppId      string
	PrivateKey string
	NotifyUrl  string
	IsPod      bool
}

type PayFailCallback func(resp interface{})

type PaySucessCallback func(resp interface{})

func (c PayClientApi) GenBuissOrderId(prix string) (ret string) {

	return fmt.Sprintf("%s_%v", prix, time.Now().UnixNano())
}
