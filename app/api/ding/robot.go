package ding

import (
	"dingtalk_server/helper"
	"github.com/gogf/gf/net/ghttp"
)

var Robot = new(rebotApi)

type rebotApi struct{}

type ReceiveReq struct {
	RobotName string `p:"robotName" v:"required#robotName字段,机器人名称不能为空"`
	MsgType   string `p:"msgType" v:"required#msgtype字段,消息类型不能为空"`
	Content   string `p:"content" v:"required#content字段,文本消息不能为空"`
	At        string `p:"at"`
}

func (c *rebotApi) Receive(r *ghttp.Request) {
	var apiReq *ReceiveReq

	if err := r.Parse(&apiReq); err != nil {
		helper.JsonExit(r, 500, "失败", err.Error())
	}

	var contextByte []byte
	//消息处理
	switch apiReq.MsgType {
	case "text":
		data, err := helper.ContextHandle(apiReq.Content)
		if err != nil {
			helper.JsonExit(r, 500, "失败", err)
		}
		contextByte = data
	case "other":
		helper.JsonExit(r, 500, "失败", "其他类型")
	}

	if len(contextByte) <= 0 {
		helper.JsonExit(r, 500, "失败", "内部异常,contextByte空指针")
	}

	msg, err := helper.SendDingMsg(string(contextByte), "wolf")
	if err != nil {
		helper.JsonExit(r, 500, "失败", err)
	}
	helper.Json(r, 200, "成功", msg)
}
