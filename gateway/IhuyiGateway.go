package gateway

import (
	"errors"
	"fmt"

	"github.com/suguer/SmsGateway/model"
	"github.com/suguer/SmsGateway/private/http"
)

type IhuyiGateway struct {
	Gateway
}

type IhuyiSendSMSMessageResponse struct {
	model.SendSMSMessageResponse
	Message string `json:"msg"`
	BizId   string `json:"smsid"`
}

func (g *IhuyiGateway) GetName() string {
	return "ihuyi"
}

func (g *IhuyiGateway) Init(c *model.Config) {
	g.Gateway.Init(c)
	g.ApiUrl = "https://106.ihuyi.com/webservice/sms.php"

}

func (g *IhuyiGateway) SendMessage(mobile *model.Phone, message *model.Message) (model.SendSMSMessageResponse, error) {
	var data IhuyiSendSMSMessageResponse
	request := http.NewHttpRequest()
	request.SetMethod("GET")
	request.SetQuery("method", "Submit")
	request.SetQuery("mobile", mobile.GetNumber())
	request.SetQuery("content", message.GetContent())
	g.buildParam(request)
	response, err := g.send(request)
	fmt.Printf("response.GetBody(): %v\n", string(response.GetBody()))
	err = model.NewSendSMSMessageResponse(&data, response.GetBody())
	if err != nil {
		return data.SendSMSMessageResponse, err
	}
	if data.Code.Val != "2" {
		return data.SendSMSMessageResponse, errors.New(data.Message)
	}
	return data.SendSMSMessageResponse, nil
}

func (g *IhuyiGateway) buildParam(request *http.HttpRequest) {
	request.SetQuery("account", g.Conf.AppID)
	request.SetQuery("password", g.Conf.AppSecret)
	request.SetQuery("format", "json")
}
