package pay

import (
	"net/http"

	confutil "black-hole.com/app/bee/bee-server/clawerutil"
	"black-hole.com/app/bee/bee-server/model"

	"github.com/iGoogle-ink/gopay"
	"github.com/skip2/go-qrcode"
)

type WeChatClientApi struct {
	PayClientApi
	MCHID           string
	APPSECRET       string
	SSLCERT_PATH    string
	CURL_PROXY_HOST string
	CURL_PROXY_PORT int
	REPORT_LEVENL   int
}

func NewWeChatClientApi(noticeurl string) (ret *WeChatClientApi) {

	baseurl := confutil.GetString("webchat", "Notify_url")

	var weclient WeChatClientApi

	weclient.AppId = confutil.GetString("webchat", "AppId")
	weclient.PrivateKey = confutil.GetString("webchat", "KEY")
	weclient.NotifyUrl = baseurl + noticeurl
	weclient.IsPod = confutil.GetBool("webchat", "IsProd")
	weclient.MCHID = confutil.GetString("webchat", "MCHID")
	weclient.APPSECRET = confutil.GetString("webchat", "APPSECRET")
	weclient.SSLCERT_PATH = confutil.GetString("webchat", "SSLCERT_PATH")
	weclient.CURL_PROXY_HOST = confutil.GetString("webchat", "CURL_PROXY_HOST")
	weclient.CURL_PROXY_PORT = confutil.GetInt("webchat", "CURL_PROXY_PORT")
	weclient.REPORT_LEVENL = confutil.GetInt("webchat", "REPORT_LEVENL")
	return &weclient
}

/**
统一下单
*/
func (c WeChatClientApi) WeChatUnifiedOrder(order *model.OrderInfo) (resp *gopay.WeChatUnifiedOrderResponse, err error) {
	//初始化微信客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := gopay.NewWeChatClient(c.AppId, c.MCHID, c.PrivateKey, c.IsPod)
	body := make(gopay.BodyMap)
	body.Set("nonce_str", gopay.GetRandomString(32))
	body.Set("body", "web测试支付")
	order.BussiesId = c.GenBuissOrderId("webchat")
	body.Set("out_trade_no", order.BussiesId)
	order.Almount = order.Price * float64(order.Quantity)
	body.Set("total_fee", int(order.Almount))
	//body.Set("spbill_create_ip", "127.0.0.1")
	body.Set("notify_url", c.NotifyUrl)
	body.Set("trade_type", gopay.TradeType_Native)
	body.Set("device_info", "WEB")
	body.Set("sign_type", gopay.SignType_MD5)

	//请求支付下单，成功后得到结果
	return client.UnifiedOrder(body)
}

func (c WeChatClientApi) WeChatCreatePCode(url string) ([]byte, error) {

	return qrcode.Encode(url, qrcode.Medium, 256)
}

func (c WeChatClientApi) WeChatPayNoticeCallBack(request *http.Request, succ PaySucessCallback, fail PayFailCallback) {

	result, err := gopay.ParseNotifyResult(request)

	if err != nil {
		fail(result)
	}
	ok, _ := gopay.VerifyPayResultSign(c.PrivateKey, gopay.SignType_MD5, result)
	if ok == false {
		fail(result)
	}
	succ(result)
}
