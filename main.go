package wxtestnotify

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	message2 "github.com/silenceper/wechat/v2/officialaccount/message"
)

func WxNotify(appId, appSecret, toUser, templateId, message string) (msgID int64, err error) {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:          appId,
		AppSecret:      appSecret,
		Token:          "",
		EncodingAESKey: "",
		Cache:          memory,
	}
	officialAccount := wc.GetOfficialAccount(cfg)
	template := officialAccount.GetTemplate()
	data := make(map[string]*message2.TemplateDataItem)
	data["date"] = &message2.TemplateDataItem{Value: message}
	return template.Send(&message2.TemplateMessage{
		ToUser:     toUser,
		TemplateID: templateId,
		Data:       data})
}
