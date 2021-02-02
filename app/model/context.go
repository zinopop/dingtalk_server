package model

//send 接口请求绑定
type DingRobotApiSendReq struct {
	RobotName string `p:"robotName" v:"required#robotName字段,机器人名称不能为空"`
	MsgType   string `p:"msgType" v:"required#msgtype字段,消息类型不能为空"`
	Content   string `p:"content" v:"required#content字段,文本消息不能为空"`
	At        string `p:"at"`
}

type DingRobotServiceSendReq struct {
	RobotName string
	MsgType   string
	Content   string
	At        string
}

/***************************************************/
