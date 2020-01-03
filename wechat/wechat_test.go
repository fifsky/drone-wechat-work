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
		Build: Build{
			Status: "success",
		},
		Url:     fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", os.Getenv("WECHAT_ROBOT_TOKEN")),
		MsgType: "markdown",
		Content: `
{{if eq .Status "success" }}
#### 🎉 ${DRONE_REPO} 构建成功
> Commit: [${DRONE_COMMIT_MESSAGE}](${DRONE_COMMIT_LINK})
> Author: ${DRONE_COMMIT_AUTHOR}
> [点击查看](${DRONE_BUILD_LINK})
{{else}}
#### ❌ ${DRONE_REPO} 构建失败
> Commit: [${DRONE_COMMIT_MESSAGE}](${DRONE_COMMIT_LINK})
> Author: ${DRONE_COMMIT_AUTHOR}
> 请立即修复!!!
> [点击查看](${DRONE_BUILD_LINK})
{{end}}
`,
	}

	err := plugin.Send()
	if err != nil {
		t.Error(err)
	}
}
