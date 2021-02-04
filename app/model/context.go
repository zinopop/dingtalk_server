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

/*===============================================================================*/
//接受钉钉消息数据绑定

type DingRobotReceiveApiReq struct {
	Msgtype           string    `p:"msgtype"`
	Text              text      `p:"text"`
	MsgID             string    `p:"msgId"`
	CreateAt          int64     `p:"createAt"`
	ConversationType  string    `p:"conversationType"`
	ConversationID    string    `p:"conversationId"`
	ConversationTitle string    `p:"conversationTitle"`
	SenderID          string    `p:"senderId"`
	SenderNick        string    `p:"senderNick"`
	SenderCorpID      string    `p:"senderCorpId"`
	SenderStaffID     string    `p:"senderStaffId"`
	ChatbotUserID     string    `p:"chatbotUserId"`
	AtUsers           []atUsers `p:"atUsers"`
}
type text struct {
	Content string `p:"content"`
}
type atUsers struct {
	DingtalkID string `p:"dingtalkId"`
	StaffID    string `p:"staffId"`
}

//=====================================
//图灵入参
type TuLingReq struct {
	Perception struct {
		InputImage struct {
			URL string `json:"url"`
		} `json:"inputImage"`
		InputText struct {
			Text string `json:"text"`
		} `json:"inputText"`
		SelfInfo struct {
			Location struct {
				City     string `json:"city"`
				Province string `json:"province"`
				Street   string `json:"street"`
			} `json:"location"`
		} `json:"selfInfo"`
	} `json:"perception"`
	ReqType  int64 `json:"reqType"`
	UserInfo struct {
		APIKey string `json:"apiKey"`
		UserID string `json:"userId"`
	} `json:"userInfo"`
}
