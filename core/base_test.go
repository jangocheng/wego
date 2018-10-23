package core_test

import (
	"github.com/godcong/wego/core"
	"testing"
)

var config *core.Config

func init() {
	cfg, _ := core.LoadConfig("D:\\workspace\\project\\goproject\\wego\\config.toml")
	config = cfg.GetSubConfig("official_account.default")
}

func TestLink(t *testing.T) {
	t.Log(core.Link("/cgi-bin/customservice/getonlinekflist"))
}

func TestBase_GetCallbackIP(t *testing.T) {
	base := core.NewBase(config)
	resp := base.GetCallbackIP()

	t.Log(resp.Error())
	t.Log(resp.ToMap())
	t.Log(string(resp.Bytes()))

}

func TestURL_ShortURL(t *testing.T) {
	resp := core.NewURL(config).ShortURL("https://mp.wechat.qq.com")
	t.Log(resp.Error())
	t.Log(resp.ToMap())
	t.Log(string(resp.Bytes()))
}
