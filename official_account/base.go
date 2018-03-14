package official_account

import "github.com/godcong/wego/core"

type Base struct {
	config  core.Config
	account *OfficialAccount
	client  *core.Client
	token   *core.AccessToken
}

func newBase(account *OfficialAccount) *Base {
	return &Base{
		config:  account.Config,
		account: account,
		client:  account.client,
		token:   account.token,
	}
}

func NewBase() *Base {
	return newBase(account)
}

//ClearQuota
//成功:
//{"errcode":0,"errmsg":"ok"}
func (b *Base) ClearQuota() core.Map {
	token := b.token.GetToken()
	params := core.Map{
		"appid": b.config.Get("app_id"),
	}
	resp := b.client.HttpPostJson(b.client.Link(CLEAR_QUOTA_URL_SUFFIX), params, core.Map{
		core.REQUEST_TYPE_QUERY.String(): token.KeyMap(),
	})
	return resp.ToMap()
}

//GetCallbackIp
//成功：
//{"ip_list":["101.226.62.77","101.226.62.78","101.226.62.79","101.226.62.80","101.226.62.81","101.226.62.82","101.226.62.83","101.226.62.84","101.226.62.85","101.226.62.86","101.226.103.59","101.226.103.60","101.226.103.61","101.226.103.62","101.226.103.63","101.226.103.69","101.226.103.70","101.226.103.71","101.226.103.72","101.226.103.73","140.207.54.73","140.207.54.74","140.207.54.75","140.207.54.76","140.207.54.77","140.207.54.78","140.207.54.79","140.207.54.80","182.254.11.203","182.254.11.202","182.254.11.201","182.254.11.200","182.254.11.199","182.254.11.198","59.37.97.100","59.37.97.101","59.37.97.102","59.37.97.103","59.37.97.104","59.37.97.105","59.37.97.106","59.37.97.107","59.37.97.108","59.37.97.109","59.37.97.110","59.37.97.111","59.37.97.112","59.37.97.113","59.37.97.114","59.37.97.115","59.37.97.116","59.37.97.117","59.37.97.118","112.90.78.158","112.90.78.159","112.90.78.160","112.90.78.161","112.90.78.162","112.90.78.163","112.90.78.164","112.90.78.165","112.90.78.166","112.90.78.167","140.207.54.19","140.207.54.76","140.207.54.77","140.207.54.78","140.207.54.79","140.207.54.80","180.163.15.149","180.163.15.151","180.163.15.152","180.163.15.153","180.163.15.154","180.163.15.155","180.163.15.156","180.163.15.157","180.163.15.158","180.163.15.159","180.163.15.160","180.163.15.161","180.163.15.162","180.163.15.163","180.163.15.164","180.163.15.165","180.163.15.166","180.163.15.167","180.163.15.168","180.163.15.169","180.163.15.170","101.226.103.0\/25","101.226.233.128\/25","58.247.206.128\/25","182.254.86.128\/25","103.7.30.21","103.7.30.64\/26","58.251.80.32\/27","183.3.234.32\/27","121.51.130.64\/27"]}
//失败:
//{"errcode":40013,"errmsg":"invalid appid"}
func (b *Base) GetCallbackIp() core.Map {
	token := b.token.GetToken()
	resp := b.client.HttpGet(b.client.Link(GETCALLBACKIP_URL_SUFFIX), core.Map{
		core.REQUEST_TYPE_QUERY.String(): token.KeyMap(),
	})
	return resp.ToMap()
}
