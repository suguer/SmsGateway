package gateway

import (
	"errors"

	"github.com/suguer/SmsGateway/model"
	"github.com/suguer/SmsGateway/private/http"
)

type JuheGateway struct {
	Gateway
}

type JuheSendSMSMessageResponse struct {
	model.SendSMSMessageResponse
	Message string `json:"msg"`
	Code    string `json:"code"`
}

func (g *JuheGateway) GetName() string {
	return "juhe"
}

func (g *JuheGateway) Init(c *model.Config) {
	g.Gateway.Init(c)
	g.ApiUrl = "http://v.juhe.cn/sms/send"

}

func (g *JuheGateway) SendMessage(mobile *model.Phone, message *model.Message) (model.SendSMSMessageResponse, error) {
	request := http.NewHttpRequest()
	request.SetQuery("mobile", mobile.GetNumber())
	request.SetQuery("tpl_id", message.GetTemplateCode())
	// request.SetQuery("vars", message.GetTemplateCode())
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

func (g *JuheGateway) buildParam(request *http.HttpRequest) {
	request.SetQuery("key", g.Conf.AppID)
}
