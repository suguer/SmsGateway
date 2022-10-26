package gateway

import (
	"encoding/base64"
	"fmt"

	"github.com/suguer/SmsGateway/model"
	"github.com/suguer/SmsGateway/private/http"
)

type LuosimaoGateway struct {
	Gateway
}
type LuosimaoSendSMSMessageResponse struct {
	model.SendSMSMessageResponse
	Message string     `json:"msg"`
	Code    model.Code `json:"error"`
}

func (g *LuosimaoGateway) Init(c *model.Config) {
	g.Gateway.Init(c)
	g.ApiUrl = "http://sms-api.luosimao.com"

}
func (g *LuosimaoGateway) GetName() string {
	return "luosimao"
}

func (g *LuosimaoGateway) SendMessage(mobile *model.Phone, message *model.Message) (model.SendSMSMessageResponse, error) {
	var data LuosimaoSendSMSMessageResponse
	request := http.NewHttpRequest()
	request.SetMethod("POST")
	request.SetPath("/v1/send.json")
	request.SetRequestBody([]byte(
		fmt.Sprintf("mobile=%s&message=%s",
			mobile.GetNumber(),
			message.GetContent()+"【"+message.GetSignName()+"】")))
	g.buildParam(request)
	response, err := g.send(request)
	if err != nil {
		return data.SendSMSMessageResponse, err
	}
	err = model.NewCommonResponse(&data, response.GetBody())
	if err != nil {
		return data.SendSMSMessageResponse, err
	}
	data.SendSMSMessageResponse.Code = data.Code
	data.SendSMSMessageResponse.Message = data.Message
	if data.SendSMSMessageResponse.Code.Val == "0" {
		data.SendSMSMessageResponse.Code.Val = "OK"
	}
	return data.SendSMSMessageResponse, nil
}

func (g *LuosimaoGateway) buildParam(request *http.HttpRequest) {
	auth := base64.StdEncoding.EncodeToString([]byte("api:" + "key-" + g.Conf.AppSecret))
	request.SetHeader("Authorization", "Basic "+auth)
}
