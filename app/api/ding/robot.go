package ding

import (
	"dingtalk_server/app/model"
	"dingtalk_server/app/service/ding"
	"dingtalk_server/helper"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Robot = new(rebotApi)

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

func (c *rebotApi) Receive(r *ghttp.Request) {
	helper.Json(r, 200, "成功")
}
