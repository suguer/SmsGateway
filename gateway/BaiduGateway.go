package gateway

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/suguer/SmsGateway/model"
	"github.com/suguer/SmsGateway/private/http"
	"github.com/suguer/SmsGateway/private/utils"
)

type BaiduGateway struct {
	Gateway
}
type BaiduSendSMSMessageResponse struct {
	model.SendSMSMessageResponse
	Code      model.Code `json:"code"`
	RequestId string     `json:"requestId"`
	Message   string     `json:"message"`
}

func (g *BaiduGateway) Init(c *model.Config) {
	g.Gateway.Init(c)
	g.ApiUrl = "http://smsv3.bj.baidubce.com"

}

func (g *BaiduGateway) GetName() string {
	return "baidu"
}

func (g *BaiduGateway) SendMessage(mobile *model.Phone, message *model.Message) (model.SendSMSMessageResponse, error) {
	request := http.NewHttpRequest()
	request.SetPath("/api/v3/sendSms")
	request.SetMethod("POST")
	args := map[string]any{
		"mobile":      "15088132466",
		"signatureId": message.GetSignName(),
		"template":    message.GetTemplateCode(),
		"contentVar":  message.GetParam(),
	}

	var data BaiduSendSMSMessageResponse
	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return data.SendSMSMessageResponse, err
	}
	request.SetRequestBody(jsonBytes)
	g.buildParam(request)
	response, err := g.send(request)
	err = model.NewCommonResponse(&data, response.GetBody())
	if err != nil {
		return data.SendSMSMessageResponse, err
	}
	data.SendSMSMessageResponse.BizId = data.BizId
	data.SendSMSMessageResponse.Code = data.Code
	data.SendSMSMessageResponse.RequestId = data.RequestId
	data.SendSMSMessageResponse.Message = data.Message
	return data.SendSMSMessageResponse, err
}

func (g *BaiduGateway) buildParam(request *http.HttpRequest) {

	format, _ := url.Parse(g.ApiUrl)
	Host := format.Host
	Timestamp := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	CanonicalURI := request.GetPath()
	CanonicalQueryString := utils.GetSignString(request.GetQueryMap(), "full")
	CanonicalHeaders := fmt.Sprintf("host:%s\nx-bce-date:%s", Host, utils.EncodeURIComponent(Timestamp))
	CanonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s", request.GetMethod(), CanonicalURI, CanonicalQueryString, CanonicalHeaders)
	authStringPrefix := fmt.Sprintf("bce-auth-v1/%s/%s/%s", g.Conf.AppID, Timestamp, "3600")
	SigningKey := utils.HMACSHA256Hex(g.Conf.AppSecret, authStringPrefix)
	Signature := utils.HMACSHA256Hex(SigningKey, CanonicalRequest)
	Authorization := fmt.Sprintf("%s/host;x-bce-date/%s", authStringPrefix, Signature)
	request.SetHeader("Authorization", Authorization)
	request.SetHeader("host", Host)
	request.SetHeader("x-bce-date", Timestamp)
}
