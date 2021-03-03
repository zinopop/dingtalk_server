package ding

import (
	"dingtalk_server/app/model"
	"dingtalk_server/app/service/ding"
	"dingtalk_server/helper"
	"dingtalk_server/helper/tuling"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// Robot 测试
var Robot = new(rebotApi)

// rebotApi 测试/
type rebotApi struct{}

//webhook发送
func (c *rebotApi) Send(r *ghttp.Request) {
	var apiReq *model.DingRobotApiSendReq
	var serviceReq *model.DingRobotServiceSendReq
	if err := r.Parse(&apiReq); err != nil {
		helper.JsonExit(r, 500, "失败", err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		helper.JsonExit(r, 500, "失败", err.Error())
	}
	data, err := ding.Robot.Send(serviceReq)
	if err != nil {
		helper.JsonExit(r, 500, "失败", err.Error())
	}
	helper.Json(r, 200, "成功", data)

}

//钉钉接收消息
func (c *rebotApi) Receive(r *ghttp.Request) {

	var apiReq *model.DingRobotReceiveApiReq
	if err := r.Parse(&apiReq); err != nil {
		helper.JsonExit(r, 500, "失败", err.Error())
	}
	re, err := tuling.Send(apiReq.Text.Content, "default")
	if err != nil {
		helper.JsonExit(r, 500, "图灵请求失败", err.Error())
	}
	j, err := gjson.DecodeToJson(re)
	if err != nil {
		helper.JsonExit(r, 500, "图灵请求返回数据解析失败", err.Error())
	}
	restring := j.GetString("results.0.values.text")

	data, err := ding.Robot.Send(&model.DingRobotServiceSendReq{
		RobotName: "wolf",
		MsgType:   "text",
		Content:   restring,
	})
	if err != nil {
		helper.JsonExit(r, 500, "失败", err.Error())
	}
	helper.Json(r, 200, "成功", data)
}
