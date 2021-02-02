package ding

import (
	"dingtalk_server/app/model"
	"dingtalk_server/helper"
	"github.com/gogf/gf/errors/gerror"
)

var Robot = new(robotService)

type robotService struct{}

func (s *robotService) Send(m *model.DingRobotServiceSendReq) (string, error) {
	var contextByte []byte
	//消息处理
	switch m.MsgType {
	case "text":
		data, err := helper.ContextHandle(m.Content)
		if err != nil {
			return "失败", err
		}
		contextByte = data
	}
	if len(contextByte) <= 0 {
		return "内部异常,contextByte空指针", gerror.New("失败")

	}
	// todo 机器人名称合法性判断
	return helper.SendDingMsg(string(contextByte), m.RobotName)
}
