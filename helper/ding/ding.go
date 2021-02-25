package ding

import (
	"dingtalk_server/helper"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/os/gcfg"
)

type MsgContext struct {
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

//
func urlBuild(names string) string {
	dingUrl := gcfg.Instance().GetString("dingTalk."+names+".sendUrl") + "?access_token=" + gcfg.Instance().GetString("dingTalk."+names+".accessToken") + "&timestamp="
	timestamp := time.Now().UnixNano() / 1e6
	timestampString := strconv.FormatInt(timestamp, 10)
	dingUrl += timestampString
	sign := timestampString + "\n" + gcfg.Instance().GetString("dingTalk."+names+".secret")
	sign = helper.ComputeHmacSha256(sign, gcfg.Instance().GetString("dingTalk."+names+".secret"))
	sign = url.QueryEscape(sign)
	dingUrl += "&sign=" + sign
	return dingUrl
}

func SendDingMsg(msg string, names string) (string, error) {
	req, err := http.NewRequest("POST", urlBuild(names), strings.NewReader(msg))
	if err != nil {
		return "生成url失败", err
	}
	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json")
	//发送请求
	resp, err := client.Do(req)

	if err != nil {
		return "请求外部出现错误", err
	}
	//关闭请求
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "读取流失败", err
	}
	return string(body), nil
}

func ContextHandle(msg string) ([]byte, error) {
	re := &MsgContext{}
	re.Text.Content = "【服务器当前时间】" + time.Now().Format("2006-01-02 15:04:05") + "\n" + msg
	re.Msgtype = "text"
	return json.Marshal(&re)
}
