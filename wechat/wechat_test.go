package wechat

import (
	"fmt"
	"os"
	"testing"
)

func TestWechatRobot_Message(t *testing.T) {
	plugin := WeChat{
		Url:     fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", os.Getenv("WECHAT_ROBOT_TOKEN")),
		MsgType: "text",
		Content: "hello",
	}
	err := plugin.Send()
	if err != nil {
		t.Error(err)
	}
}

func TestWechatRobot_MarkdownMessage(t *testing.T) {
	plugin := WeChat{
		Url:     fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", os.Getenv("WECHAT_ROBOT_TOKEN")),
		MsgType: "text",
		Content: "## 呵呵\n\n > Hello \n\n",
	}

	err := plugin.Send()
	if err != nil {
		t.Error(err)
	}
}
