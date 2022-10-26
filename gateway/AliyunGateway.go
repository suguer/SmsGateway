package gateway

import (
	"encoding/json"
	"time"

	"github.com/suguer/SmsGateway/model"
	"github.com/suguer/SmsGateway/private/http"
	"github.com/suguer/SmsGateway/private/utils"
)

type AliyunGateway struct {
	Gateway
}

func (g *AliyunGateway) Init(c *model.Config) {
	g.Gateway.Init(c)
	g.ApiUrl = "https://dysmsapi.aliyuncs.com"

}

func (g *AliyunGateway) GetName() string {
	return "aliyun"
}

func (g *AliyunGateway) SendMessage(mobile *model.Phone, message *model.Message) (model.SendSMSMessageResponse, error) {
	request := http.NewHttpRequest()
	request.SetQuery("Action", "SendSms")
	if mobile.IsChineseCode() {
		request.SetQuery("PhoneNumbers", mobile.GetNumber())
	} else {
		request.SetQuery("PhoneNumbers", mobile.GetUniversalNumber())
	}
	request.SetQuery("SignName", message.GetSignName())
	request.SetQuery("TemplateCode", message.GetTemplateCode())
	Param, err := json.Marshal(message.GetParam())
	request.SetQuery("TemplateParam", string(Param))
	g.buildParam(request)
	response, err := g.send(request)
	var data model.SendSMSMessageResponse
	err = model.NewCommonResponse(&data, response.GetBody())
	if err != nil {
		return data, err
	}
	return data, err
}
func (g *AliyunGateway) CreateSmsTemplate(template *model.Template) (model.CreateSmsTemplateResponse, error) {
	var data model.CreateSmsTemplateResponse
	request := http.NewHttpRequest()
	request.SetQuery("Action", "AddSmsTemplate")
	TemplateType := template.GetType()
	if template.GetInternational() == "1" {
		TemplateType = "4"
	}
	request.SetQuery("TemplateType", TemplateType)
	request.SetQuery("TemplateName", template.GetName())
	request.SetQuery("TemplateContent", template.GetContent())
	request.SetQuery("Remark", template.GetRemark())
	g.buildParam(request)
	response, err := g.send(request)
	if err != nil {
		return data, err
	}
	err = model.NewCommonResponse(&data, response.GetBody())
	if err != nil {
		return data, err
	}
	return data, nil
}
func (g *AliyunGateway) QuerySmsTemplate(TemplateCode string) (model.QuerySmsTemplateponse, error) {
	var data model.QuerySmsTemplateponse
	request := http.NewHttpRequest()
	request.SetQuery("Action", "QuerySmsTemplate")
	request.SetQuery("TemplateCode", TemplateCode)
	g.buildParam(request)
	response, err := g.send(request)
	if err != nil {
		return data, err
	}
	err = model.NewCommonResponse(&data, response.GetBody())
	if err != nil {
		return data, err
	}
	return data, nil
}

func (g *AliyunGateway) buildParam(request *http.HttpRequest) {
	Timestamp := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	request.SetQuery("Format", "JSON")
	request.SetQuery("Version", "2017-05-25")
	request.SetQuery("AccessKeyId", g.Conf.AppID)
	request.SetQuery("SignatureNonce", utils.GetSignatureNonce())
	request.SetQuery("Timestamp", Timestamp)
	request.SetQuery("SignatureMethod", "HMAC-SHA1")
	request.SetQuery("SignatureVersion", "1.0")
	CanonicalizedQueryString := utils.GetSignString(request.GetQueryMap(), "full")
	stringToSign := "GET" + "&" + utils.EncodeURIComponent("/") + "&" + utils.EncodeURIComponent(CanonicalizedQueryString)
	Signature := utils.HMACSHA1(g.Conf.AppSecret+"&", stringToSign)
	request.SetQuery("Signature", Signature)
}
