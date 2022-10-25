package gateway

import (
	"errors"

	"github.com/suguer/SmsGateway/model"
	"github.com/suguer/SmsGateway/private/http"
	"github.com/suguer/SmsGateway/private/utils"
)

type SmsbaoGateway struct {
	Gateway
}

type SmsbaoGatewaySendSMSMessageResponse struct {
	model.SendSMSMessageResponse
}

func (g *SmsbaoGateway) Init(c *model.Config) {
	g.Gateway.Init(c)
	g.ApiUrl = "http://api.smsbao.com"

}
func (g *SmsbaoGateway) GetName() string {
	return "smsbao"
}

func (g *SmsbaoGateway) SendMessage(mobile *model.Phone, message *model.Message) (model.SendSMSMessageResponse, error) {
	request := http.NewHttpRequest()
	request.SetPath("/sms")
	request.SetQuery("m", mobile.GetNumber())
	request.SetQuery("c", message.GetContent())
	var data model.SendSMSMessageResponse
	g.buildParam(request)
	response, err := g.send(request)
	if err != nil {
		return data, err
	}
	data.Code.Val = string(response.GetBody())
	switch data.Code.Val {
	case "-1":
		data.Message = "未知错误"
	case "30":
		data.Message = "错误密码"
	case "40":
		data.Message = "账号不存在"
	case "41":
		data.Message = "余额不足"
	case "43":
		data.Message = "IP地址限制"
	case "50":
		data.Message = "内容含有敏感词"
	case "51":
		data.Message = "手机号码不正确"
	}
	if data.Code.Val != "0" {
		return data, errors.New(data.Message)
	}
	return data, nil
}

func (g *SmsbaoGateway) buildParam(request *http.HttpRequest) {
	request.SetQuery("u", g.Conf.AppID)
	request.SetQuery("p", utils.MD5(g.Conf.AppSecret))
}
