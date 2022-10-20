package gateway

import (
	"errors"

	"github.com/suguer/SmsGateway/model"
	"github.com/suguer/SmsGateway/private/http"
)

type NowcnGateway struct {
	Gateway
}

type NowcnSendSMSMessageResponse struct {
	model.SendSMSMessageResponse
	Message string `json:"msg"`
	Code    string `json:"code"`
}

func (g *NowcnGateway) GetName() string {
	return "nowcn"
}

func (g *NowcnGateway) Init(c *model.Config) {
	g.Gateway.Init(c)
	g.ApiUrl = "http://ad1200.now.net.cn:2003/sms/sendSMS"

}

func (g *NowcnGateway) SendMessage(mobile *model.Phone, message *model.Message) (model.SendSMSMessageResponse, error) {
	request := http.NewHttpRequest()
	request.SetQuery("apiType", "3")
	request.SetQuery("mobile", mobile.GetNumber())
	request.SetQuery("content", message.GetContent())
	g.buildParam(request)
	response, err := g.send(request)

	var data NowcnSendSMSMessageResponse
	err = model.NewSendSMSMessageResponse(&data, response.GetBody())
	if err != nil {
		return data.SendSMSMessageResponse, err
	}
	if data.Code != "0" {
		return data.SendSMSMessageResponse, errors.New(data.Message)
	}
	return data.SendSMSMessageResponse, nil
}

func (g *NowcnGateway) buildParam(request *http.HttpRequest) {
	request.SetQuery("userId", g.Conf.AppID)
	request.SetQuery("password", g.Conf.AppSecret)
}
