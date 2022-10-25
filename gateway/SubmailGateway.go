package gateway

import (
	"github.com/suguer/SmsGateway/model"
	"github.com/suguer/SmsGateway/private/http"
)

type SubmailGateway struct {
	Gateway
}

type SubmailGatewaySendSMSMessageResponse struct {
	model.SendSMSMessageResponse
	BizId   string     `json:"send_id"`
	Code    model.Code `json:"code"`
	Message string     `json:"msg"`
}

func (g *SubmailGateway) Init(c *model.Config) {
	g.Gateway.Init(c)
	g.ApiUrl = "https://api-v4.mysubmail.com/"

}

func (g *SubmailGateway) GetName() string {
	return "submail"
}
func (g *SubmailGateway) SendMessage(mobile *model.Phone, message *model.Message) (model.SendSMSMessageResponse, error) {
	request := http.NewHttpRequest()
	var data SubmailGatewaySendSMSMessageResponse
	if len(message.GetTemplateCode()) > 0 {
		//发送模板类短信
		request.SetPath("/sms/xsend")
	} else {
		//发送内容
		request.SetPath("/sms/send")
	}
	g.buildParam(request)
	return data.SendSMSMessageResponse, nil
}

func (g *SubmailGateway) buildParam(request *http.HttpRequest) {
}
