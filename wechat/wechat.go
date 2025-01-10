package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"text/template"
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

	Build struct {
		Owner   string
		Name    string
		Tag     string
		Event   string
		Number  int
		Commit  string
		Ref     string
		Branch  string
		Author  string
		Message string
		Status  string
		Link    string
		Started int64
		Created int64
	}

	Response struct {
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}

	WeChat struct {
		Build   Build
		Urls    []string
		MsgType string
		ToUser  string
		Content string
	}
)

func jsonEncode(d interface{}) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(buf)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(d)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

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

	buf, err := jsonEncode(we)
	if err != nil {
		return err
	}

	return c.call(buf)
}

func (c *WeChat) Template(temp string) ([]byte, error) {
	tmpl, err := template.New("wechat").Parse(temp)
	if err != nil {
		return nil, fmt.Errorf("template parse error %w %s", err, temp)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, c.Build)
	if err != nil {
		return nil, fmt.Errorf("template execute error %w", err)
	}

	return buf.Bytes(), nil
}

func (c *WeChat) call(buf *bytes.Buffer) error {
	var errs []string

	for _, url := range c.Urls {
		resp, err := c.postJson(url, buf)
		if err != nil {
			errs = append(errs, fmt.Sprintf("request to %s failed: %v", url, err))
			continue
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			errs = append(errs, fmt.Sprintf("reading response from %s failed: %v", url, err))
			continue
		}

		ret := &Response{}
		if err := json.Unmarshal(data, ret); err != nil {
			errs = append(errs, fmt.Sprintf("parsing response from %s failed: %v", url, err))
			continue
		}

		if ret.Errcode != 0 {
			errs = append(errs, fmt.Sprintf("push to %s failed: %s [%d]", url, ret.Errmsg, ret.Errcode))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("encountered following errors while sending messages:\n%s", strings.Join(errs, "\n"))
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

	buf, err := jsonEncode(we)
	if err != nil {
		return err
	}

	return c.call(buf)
}

func (c *WeChat) postJson(url string, body *bytes.Buffer) (*http.Response, error) {
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

	tempBuf, err := c.Template(c.Content)
	if err != nil {
		return err
	}

	if c.MsgType == "text" {
		return c.Message(strings.TrimSpace(string(tempBuf)), at...)
	}

	if c.MsgType == "markdown" {
		return c.MarkdownMessage(strings.TrimSpace(string(tempBuf)), at...)
	}

	return fmt.Errorf("no support msgtype %s", c.MsgType)
}
