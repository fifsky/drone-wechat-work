package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type (
	Request struct {
		Msgtype string                 `json:"msgtype"`
		Text    map[string]interface{} `json:"text"`
	}

	MarkdownRequest struct {
		Msgtype  string                 `json:"msgtype"`
		Markdown map[string]interface{} `json:"markdown"`
	}

	Response struct {
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}

	WeChat struct {
		Url     string
		MsgType string
		ToUser  string
		Content string
	}
)

func (c *WeChat) MarkdownMessage(md string, at ...string) error {
	we := &MarkdownRequest{
		Msgtype: "markdown",
		Markdown: map[string]interface{}{
			"content": md,
		},
	}

	if len(at) > 0 {
		we.Markdown["mentioned_mobile_list"] = at
	}

	buf, err := json.Marshal(we)
	if err != nil {
		return err
	}

	return c.call(buf)
}

func (c *WeChat) call(buf []byte) error {
	resp, err := c.postJson(c.Url, buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if data, err := ioutil.ReadAll(resp.Body); err == nil {
		ret := &Response{}
		err := json.Unmarshal(data, ret)
		if err != nil {
			return err
		}

		if ret.Errcode != 0 {
			return errors.New("ding response error:" + ret.Errmsg + "[" + strconv.Itoa(ret.Errcode) + "]")
		}
	}

	return nil
}

func (c *WeChat) Message(content string, at ...string) error {
	we := &Request{
		Msgtype: "text",
		Text: map[string]interface{}{
			"content": content,
		},
	}

	if len(at) > 0 {
		we.Text["mentioned_mobile_list"] = at
	}

	buf, err := json.Marshal(we)
	if err != nil {
		return err
	}

	return c.call(buf)
}

func (c *WeChat) postJson(url string, data []byte) (*http.Response, error) {
	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	client := &http.Client{}
	client.Timeout = 5 * time.Second

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *WeChat) Send() error {
	var at []string
	if c.ToUser != "" {
		at = strings.Split(c.ToUser, ",")
	}

	if c.MsgType == "text" {
		return c.Message(c.Content, at...)
	}

	if c.MsgType == "markdown" {
		return c.MarkdownMessage(c.Content, at...)
	}

	return fmt.Errorf("no support msgtype %s", c.MsgType)
}
