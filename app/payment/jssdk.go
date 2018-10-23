package payment

import (
	"strings"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*JSSDK JSSDK */
type JSSDK struct {
	*Payment
}

func newJSSDK(p *Payment) *JSSDK {
	return &JSSDK{
		Payment: p,
	}
}

/*NewJSSDK NewJSSDK */
func NewJSSDK(config *core.Config) *JSSDK {
	return newJSSDK(NewPayment(config))
}

func (j *JSSDK) getURL() string {
	return core.GetServerIP()
}

/*BridgeConfig bridge 设置 */
func (j *JSSDK) BridgeConfig(pid string) util.Map {
	appid := j.config.Get("sub_appid")
	if appid == "" {
		appid = j.config.Get("app_id")
	}

	m := util.Map{
		"appId":     appid,
		"timeStamp": util.Time(),
		"nonceStr":  util.GenerateNonceStr(),
		"package":   strings.Join([]string{"prepay_id", pid}, "="),
		"signType":  "MD5",
	}

	m.Set("paySign", core.GenerateSignature(m, j.config.GetString("key"), core.MakeSignMD5))

	return m
}

/*SdkConfig sdk 设置 */
func (j *JSSDK) SdkConfig(pid string) util.Map {
	config := j.BridgeConfig(pid)

	config.Set("timestamp", config.Get("timeStamp"))
	config.Delete("timeStamp")

	return config
}

/*AppConfig app 设置 */
func (j *JSSDK) AppConfig(pid string) util.Map {
	m := util.Map{
		"appid":     j.config.Get("app_id"),
		"partnerid": j.config.Get("mch_id"),
		"prepayid":  pid,
		"noncestr":  util.GenerateNonceStr(),
		"timestamp": util.Time(),
		"package":   "Sign=WXPay",
	}

	m.Set("sign", core.GenerateSignature(m, j.config.GetString("aes_key"), core.MakeSignMD5))
	return m
}

/*ShareAddressConfig 共享地址设置 */
func (j *JSSDK) ShareAddressConfig(accessToken interface{}) util.Map {
	token := ""
	switch v := accessToken.(type) {
	case core.AccessToken:
		token0 := v.GetToken()
		token = token0.ToJSON()
	case string:
		token = accessToken.(string)
	}
	m := util.Map{
		"appId":     j.config.Get("app_id"),
		"scope":     "jsapi_address",
		"timeStamp": util.Time(),
		"nonceStr":  util.GenerateNonceStr(),
		"signType":  "SHA1",
	}

	sm := util.Map{
		"appid":       m.Get("appId"),
		"url":         j.getURL(),
		"timestamp":   m.Get("timeStamp"),
		"noncestr":    m.Get("nonceStr"),
		"accesstoken": token,
	}

	m.Set("addrSign", util.SHA1(sm.URLEncode()))

	return m
}
