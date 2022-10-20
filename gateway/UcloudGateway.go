package gateway

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"

	"github.com/suguer/SmsGateway/model"
	"github.com/suguer/SmsGateway/private/http"
	"github.com/suguer/SmsGateway/private/utils"
)

type UcloudGateway struct {
	Gateway
}

type UcloudSendSMSMessageResponse struct {
	model.SendSMSMessageResponse
	BizId     string `json:"SessionNo"`
	Code      int    `json:"RetCode"`
	RequestId string `json:"ReqUuid"`
}

func (g *UcloudGateway) Init(c *model.Config) {
	g.Gateway.Init(c)
	g.ApiUrl = "https://api.ucloud.cn/"

}

func (g *UcloudGateway) GetName() string {
	return "ucloud" + g.Conf.AppID
}

func (g *UcloudGateway) SendMessage(mobile *model.Phone, message *model.Message) (model.SendSMSMessageResponse, error) {
	request := http.NewHttpRequest()
	request.SetQuery("Action", "SendUSMSMessage")
	request.SetQuery("PhoneNumbers.0", mobile.GetNumber())
	request.SetQuery("TemplateId", message.GetTemplateCode())
	request.SetQuery("SigContent", message.GetSignName())
	for k, v := range message.GetParam() {
		request.SetQuery("TemplateParams."+k, v)
	}
	g.buildParam(request)
	response, err := g.send(request)
	var data UcloudSendSMSMessageResponse
	if err != nil {
		return data.SendSMSMessageResponse, err
	}
	err = model.NewSendSMSMessageResponse(&data, response.GetBody())
	if err != nil {
		return data.SendSMSMessageResponse, err
	}
	if data.Code != 0 {
		return data.SendSMSMessageResponse, errors.New(data.Message)
	}
	return data.SendSMSMessageResponse, nil
}

func (g *UcloudGateway) buildParam(request *http.HttpRequest) {
	request.SetQuery("PublicKey", g.Conf.AppID)
	SignString := utils.GetSignString(request.GetQueryMap(), "") + g.Conf.AppSecret
	t := sha1.New()
	io.WriteString(t, SignString)
	Signature := fmt.Sprintf("%x", t.Sum(nil))
	request.SetQuery("Signature", Signature)
}
