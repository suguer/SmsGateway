package gateway

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/suguer/SmsGateway/model"
	"github.com/suguer/SmsGateway/private/http"
	"github.com/suguer/SmsGateway/private/utils"
)

type QiniuGateway struct {
	Gateway
}

type QiniuSendSMSMessageResponse struct {
	model.SendSMSMessageResponse
	Message string `json:"message"`
	Code    string `json:"error"`
	BizId   string `json:"message_id"`
}

func (g *QiniuGateway) GetName() string {
	return "qiniu"
}

func (g *QiniuGateway) Init(c *model.Config) {
	g.Gateway.Init(c)
	g.ApiUrl = "https://sms.qiniuapi.com"

}

func (g *QiniuGateway) SendMessage(mobile *model.Phone, message *model.Message) (model.SendSMSMessageResponse, error) {
	var data QiniuSendSMSMessageResponse
	request := http.NewHttpRequest()
	request.SetMethod("POST")
	request.SetPath("/v1/message/single")
	args := map[string]any{
		"mobile":      mobile.GetNumber(),
		"template_id": message.GetTemplateCode(),
		"parameters":  message.GetParam(),
	}
	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return data.SendSMSMessageResponse, err
	}
	request.SetRequestBody(jsonBytes)
	g.buildParam(request)
	response, err := g.send(request)

	err = model.NewSendSMSMessageResponse(&data, response.GetBody())
	if err != nil {
		return data.SendSMSMessageResponse, err
	}
	if len(data.Code) != 0 {
		return data.SendSMSMessageResponse, errors.New(data.Message)
	}
	return data.SendSMSMessageResponse, nil
}

func (g *QiniuGateway) buildParam(request *http.HttpRequest) {
	format, _ := url.Parse(g.ApiUrl)
	Host := format.Host
	data := request.GetMethod() + " " + request.GetPath()
	data += "\nHost: " + Host
	data += "\nContent-Type: " + http.MimeJSON
	data += "\n\n"
	if len(request.GetRequestBody()) > 0 {
		data += string(request.GetRequestBody())
	}
	encodedSign := utils.HMACSHA1(g.Conf.AppSecret, data)
	request.SetHeader("Host", Host)
	request.SetHeader(http.HeaderNameContentType, http.MimeJSON)
	request.SetHeader("Authorization", fmt.Sprintf("Qiniu %s:%s", g.Conf.AppID, encodedSign))
}
