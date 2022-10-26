package gateway

import (
	"github.com/suguer/SmsGateway/model"
	"github.com/suguer/SmsGateway/private/http"
)

type JuheGateway struct {
	Gateway
}

type JuheSendSMSMessageResponse struct {
	model.SendSMSMessageResponse
	Message string     `json:"msg"`
	Code    model.Code `json:"error_code"`
}

func (g *JuheGateway) GetName() string {
	return "juhe"
}

func (g *JuheGateway) Init(c *model.Config) {
	g.Gateway.Init(c)
	g.ApiUrl = "http://v.juhe.cn/"

}

func (g *JuheGateway) SendMessage(mobile *model.Phone, message *model.Message) (model.SendSMSMessageResponse, error) {
	request := http.NewHttpRequest()
	request.SetPath("sms/send")
	request.SetQuery("mobile", mobile.GetNumber())
	request.SetQuery("tpl_id", message.GetTemplateCode())
	g.buildParam(request)
	response, err := g.send(request)

	var data JuheSendSMSMessageResponse
	err = model.NewCommonResponse(&data, response.GetBody())
	if err != nil {
		return data.SendSMSMessageResponse, err
	}
	data.SendSMSMessageResponse.BizId = data.BizId
	data.SendSMSMessageResponse.Code = data.Code
	data.SendSMSMessageResponse.RequestId = data.RequestId
	data.SendSMSMessageResponse.Message = data.Message
	if data.SendSMSMessageResponse.Code.Val == "0" {
		data.SendSMSMessageResponse.Code.Val = "OK"
	}
	return data.SendSMSMessageResponse, nil
}

func (g *JuheGateway) buildParam(request *http.HttpRequest) {
	request.SetQuery("key", g.Conf.AppID)
}
