# drone-wechat-work

wechat work robot plugin for drone

### Usage

```yaml
pipeline:
  wechat:
    image: fifsky/drone-wechat-work
    url: https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxxx
    content: "Build Number: ${DRONE_BUILD_NUMBER} failed. ${DRONE_COMMIT_AUTHOR} please fix. Check the results here: ${DRONE_BUILD_LINK} "
    msgtype: "text"
    touser: "13812345678,13898754321"
    when:
      status: [ failure ]
```



## Options

| option | type | required | default | description |
| --- | --- | --- | --- | --- |
| url | string | Yes | none | The full address of webhook: https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxxx |
| type | string | No | text | message type，support (text,markdown) |
| content | string | Yes | none |  Message content, text or markdown or json string |
| touser | string | No | none | At user,Use commas to separate, for example: 13812345678,13898754321 or all |
