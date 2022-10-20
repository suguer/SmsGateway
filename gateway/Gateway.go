package gateway

import (
	"errors"

	"github.com/suguer/SmsGateway/model"
	"github.com/suguer/SmsGateway/private/http"
)

type GatewayInterface interface {
	Init(c *model.Config)

	GetName() string

	// SendMessage 发送短信 ,Struct,错误码
	SendMessage(mobile *model.Phone, message *model.Message) (model.SendSMSMessageResponse, error)

	SendBatchSms(mobile []*model.Phone, message *model.Message) (model.SendSMSMessageResponse, error)

	CreateSmsTemplate(template *model.Template) (model.CreateSmsTemplateResponse, error)

	QuerySmsTemplate(TemplateCode string) (model.CommonResponse, error)
}

type Gateway struct {
	Conf       *model.Config
	httpClient http.HttpClient
	ApiUrl     string
}

func (g *Gateway) SendMessage(mobile *model.Phone, message *model.Message) (model.SendSMSMessageResponse, error) {
	var response model.SendSMSMessageResponse
	return response, errors.New("unsupport")
}

func (g *Gateway) SendBatchSms(mobile []*model.Phone, message *model.Message) (model.SendSMSMessageResponse, error) {
	var response model.SendSMSMessageResponse
	return response, errors.New("unsupport")
}

func (g *Gateway) CreateSmsTemplate(template *model.Template) (model.CreateSmsTemplateResponse, error) {
	var response model.CreateSmsTemplateResponse
	return response, errors.New("unsupport")
}

func (g *Gateway) QuerySmsTemplate(TemplateCode string) (model.CommonResponse, error) {
	var response model.CommonResponse
	return response, errors.New("unsupport")
}

func (g *Gateway) buildParam(request *http.HttpRequest) {
}

func (g *Gateway) send(request *http.HttpRequest, Action ...string) (*http.HttpResponse, error) {
	request.SetURL(g.ApiUrl + request.GetPath())
	response, err := g.httpClient.Send(request)
	return response, err
}

func (g *Gateway) Init(c *model.Config) {
	g.Conf = c
	g.httpClient = http.NewHttpClient()
}

func NewGatewayInterface(platform string, c *model.Config) (GatewayInterface, error) {
	var gateway GatewayInterface
	switch platform {
	case "ucloud":
		gateway = &UcloudGateway{}
	case "aliyun":
		gateway = &AliyunGateway{}
	case "nowcn":
		gateway = &NowcnGateway{}
	case "tencent":
		gateway = &TencentGateway{}
	case "baidu":
		gateway = &BaiduGateway{}
	case "qiniu":
		gateway = &QiniuGateway{}
	case "ihuiyi":
		gateway = &IhuyiGateway{}
	default:
		return nil, errors.New("unspport")
	}
	gateway.Init(c)
	return gateway, nil

}
