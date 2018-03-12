package core

import (
	"github.com/godcong/wego/core/message"
)

type Button Map

//{

//	Type     string //是	菜单的响应动作类型，view表示网页类型，click表示点击类型，miniprogram表示小程序类型
//	Name     string //是	菜单标题，不超过16个字节，子菜单不超过60个字节
//	Key      string //click等点击类型必须	菜单KEY值，用于消息接口推送，不超过128字节
//	Url      string //view、miniprogram类型必须	网页 链接，用户点击菜单可打开链接，不超过1024字节。 type为miniprogram时，不支持小程序的老版本客户端将打开本url。
//	MediaId  string //media_id类型和view_limited类型必须	调用新增永久素材接口返回的合法media_id
//	AppId    string //miniprogram类型必须	小程序的appid（仅认证公众号可配置）
//	Pagepath string //string miniprogram类型必须	小程序的页面路径
//}

func NewClickButton(name, key string) *Button {
	return newButton(message.EventClick.String(), Map{"name": name, "key": key})

}

func NewViewButton(name, url string) *Button {
	return newButton(message.EventView.String(), Map{"name": name, "url": url})
}

func NewScanCodeWaitMsgButton(name, key string) *Button {
	return newButton(message.EventScancodeWaitmsg.String(), Map{"name": name, "key": key})
}

func NewScanCodePushButton(name, key string) *Button {
	return newButton(message.EventScancodePush.String(), Map{"name": name, "key": key})
}

func NewPicSysPhotoButton(name, key string) *Button {
	return newButton(message.EventPicSysphoto.String(), Map{"name": name, "key": key})
}

func NewPicPhotoOrAlbumButton(name, key string) *Button {
	return newButton(message.EventPicPhotoOrAlbum.String(), Map{"name": name, "key": key})
}

func NewPicWeixinButton(name, key string) *Button {
	return newButton(message.EventPicWeixin.String(), Map{"name": name, "key": key})
}

func NewMediaIDButton(name, mediaId string) *Button {
	return newButton("media_id", Map{"name": name, "media_id": mediaId})
}

func NewViewLimitedButton(name, mediaId string) *Button {
	return newButton("view_limited", Map{"name": name, "media_id": mediaId})
}

func NewMiniProgramButton(name, url, pagepath string) *Button {
	return newButton("miniprogram", Map{"name": name, "url": url, "pagepath": pagepath})
}

func NewSubButton(name string, sub []*Button) *Button {
	return newButton("", Map{"name": name, "sub_button": sub})
}

func newButton(typ string, val Map) *Button {
	v := make(Button)
	if typ != "" {
		(*Map)(&v).Set("type", typ)
	}
	(*Map)(&v).Join(val)
	return &v
}

func (b *Button) SetSub(name string, sub []*Button) *Button {
	*b = make(Button)
	(*Map)(b).Set("name", name)
	(*Map)(b).Set("sub_button", sub)
	return b
}