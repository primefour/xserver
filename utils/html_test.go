package utils

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"testing"
	"time"
)

func TestHtml(t *testing.T) {
	InitTranslationsWithDir("i18n")
	InitHTMLWithDir("templates")

	nowSec := time.Now().Unix()
	tm := time.Unix(nowSec, 0)
	timezone, _ := tm.Zone()

	bodyPage := NewHTMLTemplate("post_body", "zh-CN")
	bodyPage.Props["SiteURL"] = "www.fpbbc.com"
	bodyPage.Props["PostMessage"] = "Hello World Post Message"
	bodyPage.Props["TeamLink"] = "Team link"
	bodyPage.Props["BodyText"] = "body test message"
	bodyPage.Props["Button"] = T("api.templates.post_body.button")
	bodyPage.Html["Info"] = template.HTML(T("api.templates.post_body.info",
		map[string]interface{}{"ChannelName": "channelName", "SenderName": "senderName",
			"Hour": fmt.Sprintf("%02d", tm.Hour()), "Minute": fmt.Sprintf("%02d", tm.Minute()),
			"TimeZone": timezone, "Month": tm.Month(), "Day": tm.Day()}))
	xbuf := bodyPage.Render()
	xbytes := []byte(xbuf)
	ioutil.WriteFile("/home/crazyhorse/CodeWork/GoWorkSpace/case/src/github.com/primefour/xserver/test.html", xbytes, 0664)
}
