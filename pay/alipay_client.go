package pay

import (
	"net/http"

	// "black-hole.com/pkg/gopay"
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/alipay"
)

type AlipayClientApi struct {
	PayClientApi
	PublKey  string
	Charset  string
	SignType string
}

/**
回调异步查询订单
*/
func (c *AlipayClientApi) TradeQuery(bussid string) (resp *alipay.TradeQueryResponse, err error) {
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := alipay.NewClient(c.AppId, c.PrivateKey, c.IsPod)
	//配置公共参数
	client.SetCharset(c.Charset).
		SetSignType(c.SignType).
		SetNotifyUrl(c.NotifyUrl)

	////请求参数
	body := make(gopay.BodyMap)
	body.Set("out_trade_no", bussid)

	//创建订单
	return client.TradeQuery(body)
}

/**
回调验签
*/
func (c *AlipayClientApi) VerifySign(req *http.Request) (bool, *alipay.NotifyRequest) {
	ret, er1 := alipay.ParseNotifyResult(req)
	if er1 != nil {
		return false, nil
	}
	if ok, err := alipay.VerifySign(c.PublKey, ret); err == nil {
		return ok, ret
	} else {
		return false, ret
	}
}

func (c *AlipayClientApi) NoticeCallBak(request *http.Request, succ PaySucessCallback, fail PayFailCallback) {

	if ok, noticresp := c.VerifySign(request); ok == false {
		fail(noticresp)
		return
	} else {
		if noticresp.TradeStatus == "TRADE_SUCCESS" ||
			noticresp.TradeStatus == "TRADE_FINISHED" { //支付成功
			succ(noticresp)
		} else {
			fail(noticresp)
		}
	}
}
