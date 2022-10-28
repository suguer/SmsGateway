package gateway

import (
	"strconv"

	"github.com/suguer/SmsGateway/model"
	"github.com/suguer/SmsGateway/private/http"
)

type TinreeGateway struct {
	Gateway
}

type TinreeSendSMSMessageResponse struct {
	model.SendSMSMessageResponse
	BizId   string `json:"smUuid"`
	Message string `json:"msg"`
}
type TinreeCreateSmsTemplateResponse struct {
	model.CreateSmsTemplateResponse
	Code    model.Code `json:"ret"`
	Message string     `json:"msg"`
	Data    struct {
		TemplateCode int `json:"templateId"`
	}
}
type TinreeQuerySmsTemplateponse struct {
	model.QuerySmsTemplateponse
	Message string `json:"msg"`
	Data    struct {
		ApplyStatus int `json:"applyStatus"`
	}
}

func (g *TinreeGateway) Init(c *model.Config) {
	g.Gateway.Init(c)
	g.ApiUrl = "http://api.tinree.com"

}
func (g *TinreeGateway) GetName() string {
	return "tinree"
}

func (g *TinreeGateway) SendMessage(mobile *model.Phone, message *model.Message) (model.SendSMSMessageResponse, error) {
	request := http.NewHttpRequest()
	request.SetPath("/api/v2/single_send")
	request.SetQuery("mobile", mobile.GetNumber())
	request.SetQuery("sign", "【"+message.GetSignName()+"】")
	request.SetQuery("content", message.GetContent())
	request.SetQuery("templateId", message.GetTemplateCode())
	g.buildParam(request)
	response, err := g.send(request)
	var data TinreeSendSMSMessageResponse
	if err != nil {
		return data.SendSMSMessageResponse, err
	}
	err = model.NewCommonResponse(&data, response.GetBody())
	if err != nil {
		return data.SendSMSMessageResponse, err
	}
	data.SendSMSMessageResponse.BizId = data.BizId
	data.SendSMSMessageResponse.Code = data.Code
	data.SendSMSMessageResponse.RequestId = data.RequestId
	data.SendSMSMessageResponse.Message = data.Message
	if data.Code.Val == "0" {
		data.SendSMSMessageResponse.Code.Val = "OK"
	}
	return data.SendSMSMessageResponse, nil
}

func (g *TinreeGateway) CreateSmsTemplate(template *model.Template) (model.CreateSmsTemplateResponse, error) {
	var data TinreeCreateSmsTemplateResponse
	request := http.NewHttpRequest()
	request.SetPath("/open/api/addTemplate")
	request.SetQuery("templateName", template.GetName())
	request.SetQuery("template", template.GetContent())
	request.SetQuery("description", template.GetRemark())
	categoryId := "1"
	if template.GetType() == "1" {
		categoryId = "2"
	} else if template.GetType() == "2" {
		categoryId = "3"
	}
	request.SetQuery("categoryId", categoryId)
	g.buildParam(request)
	response, err := g.send(request)
	if err != nil {
		return data.CreateSmsTemplateResponse, err
	}
	err = model.NewCommonResponse(&data, response.GetBody())
	if err != nil {
		return data.CreateSmsTemplateResponse, err
	}
	data.CreateSmsTemplateResponse.TemplateCode = strconv.Itoa(data.Data.TemplateCode)
	return data.CreateSmsTemplateResponse, nil
}

func (g *TinreeGateway) QuerySmsTemplate(TemplateCode string) (model.QuerySmsTemplateponse, error) {
	var data TinreeQuerySmsTemplateponse
	request := http.NewHttpRequest()
	request.SetPath("/query/getTemplate")
	request.SetQuery("templateId", TemplateCode)
	g.buildParam(request)
	response, err := g.send(request)
	if err != nil {
		return data.QuerySmsTemplateponse, err
	}
	err = model.NewCommonResponse(&data, response.GetBody())
	if err != nil {
		return data.QuerySmsTemplateponse, err
	}
	data.TemplateStatus = data.Data.ApplyStatus - 1
	return data.QuerySmsTemplateponse, nil
}

func (g *TinreeGateway) buildParam(request *http.HttpRequest) {
	request.SetQuery("accesskey", g.Conf.AppID)
	request.SetQuery("secret", g.Conf.AppSecret)
}
