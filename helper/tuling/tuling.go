package tuling

import (
	"dingtalk_server/app/model"
	"encoding/json"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcfg"
)

func Send(content string, tuNames string) (string, error) {
	params := model.TuLingReq{}
	params.Perception.InputText.Text = content
	params.UserInfo.APIKey = gcfg.Instance().GetString("tuLing." + tuNames + ".apiKey")
	params.UserInfo.UserID = gcfg.Instance().GetString("tuLing." + tuNames + ".userId")
	resultByte, err := json.Marshal(&params)
	if err != nil {
		return "转换失败", err
	}
	result, err := g.Client().Post(gcfg.Instance().GetString("tuLing."+tuNames+".sendUrl"), string(resultByte))
	if err != nil {
		return "请求失败", err
	}
	return result.ReadAllString(), nil
}
