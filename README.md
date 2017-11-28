# drone-wechat (WORK IN PROGRESS)

Drone plugin to send build status notifications via WeChat for Work. For usage
information please look at [the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
go build
```

## CHANGE BELOW

## Docker

Build the Docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-wechat
docker build --rm -t plugins/drone-wechat .
```

### Usage

```
docker run --rm \
  -e PLUGIN_CORPID=corpid \
  -e PLUGIN_CORP_SECRET=corpsecret \
  -e PLUGIN_AGENT_ID=agentid \
  -e PLUGIN_DEBUG=true \
  -e PLUGIN_MSG_URL=url \
  -e PLUGIN_BTN_TXT=true \
  -e PLUGIN_TITLE=title \
  -e PLUGIN_DESCRIPTION=description \
  -e DRONE_REPO_OWNER=octocat \
  -e DRONE_REPO_NAME=hello-world \
  -e DRONE_COMMIT_SHA=7fd1a60b01f91b314f59955a4e4d4e80d8edf11d \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_COMMIT_AUTHOR=octocat \
  -e DRONE_BUILD_NUMBER=1 \
  -e DRONE_BUILD_STATUS=success \
  -e DRONE_BUILD_LINK=http://github.com/octocat/hello-world \
  -e DRONE_TAG=1.0.0 \
  plugins/drone-wechat
```
