# drone-wechat-work
[![Build and Test](https://github.com/fifsky/drone-wechat-work/actions/workflows/build.yml/badge.svg)](https://github.com/fifsky/drone-wechat-work/actions/workflows/build.yml)

wechat work robot plugin for drone

### Usage

```yaml
  - name: notify
    image: fifsky/drone-wechat-work
    pull: always
    settings:
      url: https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=XXX-XXXX-XXX-XXXXX
      msgtype: markdown
      content: |
        {{if eq .Status "success" }}
        #### 🎉 ${DRONE_REPO} [构建成功](${DRONE_BUILD_LINK})
        CommitID: [${DRONE_COMMIT_SHA:0:8}](${DRONE_COMMIT_LINK})
        Author: ${DRONE_COMMIT_AUTHOR}
        {{ .Message }}
        {{else}}
        #### ❌ ${DRONE_REPO} [构建失败](${DRONE_BUILD_LINK})
        CommitID: [${DRONE_COMMIT_SHA:0:8}](${DRONE_COMMIT_LINK})
        Author: ${DRONE_COMMIT_AUTHOR}
        Failed Steps: ${DRONE_FAILED_STEPS}
        {{ .Message }}
        {{end}}
    when:
      status:
        - failure
        - success
```



## Options

| option | type | required | default | description |
| --- | --- | --- | --- | --- |
| url | []string | Yes | none | The full address of webhook: https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxxx |
| type | string | No | text | message type，support (text,markdown) |
| content | string | Yes | none |  Message content, text or markdown or json string |
| touser | string | No | none | At user,Use commas to separate, for example: 13812345678,13898754321 or all |
