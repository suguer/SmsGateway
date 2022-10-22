package gateway

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/suguer/SmsGateway/model"
	"github.com/suguer/SmsGateway/private/http"
	"github.com/suguer/SmsGateway/private/utils"
)

type UcloudGateway struct {
	Gateway
}

type UcloudSendSMSMessageResponse struct {
	model.SendSMSMessageResponse
	BizId     string     `json:"SessionNo"`
	Code      model.Code `json:"RetCode"`
	RequestId string     `json:"ReqUuid"`
}
type UcloudCreateSmsTemplateResponse struct {
	model.CreateSmsTemplateResponse
	Code         model.Code `json:"RetCode"`
	TemplateCode string     `json:"TemplateId"`
	RequestId    string     `json:"ReqUuid"`
}

type UcloudQuerySmsTemplateponse struct {
	model.QuerySmsTemplateponse
	Code      model.Code `json:"RetCode"`
	RequestId string     `json:"ReqUuid"`
	Data      struct {
		TemplateCode string `json:"TemplateId"`
		Status       int
		ErrDesc      string
	}
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
	err = model.NewCommonResponse(&data, response.GetBody())
	if err != nil {
		return data.SendSMSMessageResponse, err
	}
	data.SendSMSMessageResponse.BizId = data.BizId
	data.SendSMSMessageResponse.Code = data.Code
	data.SendSMSMessageResponse.RequestId = data.RequestId
	data.SendSMSMessageResponse.Message = data.Message
	if data.Code.Val != "0" {
		return data.SendSMSMessageResponse, errors.New(data.Message)
	}
	return data.SendSMSMessageResponse, nil
}

func (g *UcloudGateway) CreateSmsTemplate(template *model.Template) (model.CreateSmsTemplateResponse, error) {
	var data UcloudCreateSmsTemplateResponse
	request := http.NewHttpRequest()
	request.SetQuery("Action", "CreateUSMSTemplate")
	International := "false"
	if template.GetInternational() == "1" {
		International = "true"
	}
	request.SetQuery("International", International)
	request.SetQuery("TemplateName", template.GetName())
	request.SetQuery("Template", template.GetContent())
	request.SetQuery("Remark", template.GetRemark())
	request.SetQuery("Purpose", strconv.Itoa(template.GetType2Int()+1))
	if request.GetQuery("Purpose") == "3" {
		request.SetQuery("UnsubscribeInfo", "退订回T")
	}
	g.buildParam(request)
	response, err := g.send(request)
	if err != nil {
		return data.CreateSmsTemplateResponse, err
	}
	err = model.NewCommonResponse(&data, response.GetBody())
	if err != nil {
		return data.CreateSmsTemplateResponse, err
	}
	data.CreateSmsTemplateResponse.TemplateCode = data.TemplateCode
	data.CreateSmsTemplateResponse.Code = data.Code
	data.CreateSmsTemplateResponse.RequestId = data.RequestId
	if data.Code.Val != "0" {
		return data.CreateSmsTemplateResponse, errors.New(data.Message)
	}
	return data.CreateSmsTemplateResponse, nil
}

func (g *UcloudGateway) QuerySmsTemplate(TemplateCode string) (model.QuerySmsTemplateponse, error) {
	var data UcloudQuerySmsTemplateponse
	request := http.NewHttpRequest()
	request.SetQuery("Action", "QueryUSMSTemplate")
	request.SetQuery("TemplateId", TemplateCode)
	g.buildParam(request)
	response, err := g.send(request)
	if err != nil {
		return data.QuerySmsTemplateponse, err
	}
	err = model.NewCommonResponse(&data, response.GetBody())
	fmt.Printf("string(response.GetBody()): %v\n", string(response.GetBody()))
	if err != nil {
		return data.QuerySmsTemplateponse, err
	}
	//短信模板状态；状态说明：0-待审核，1-审核中，2-审核通过，3-审核未通过，4-被禁用
	StatusMap := map[int]int{0: 3, 1: 0, 2: 1, 3: 2, 4: 4}
	data.QuerySmsTemplateponse.TemplateStatus = StatusMap[data.Data.Status]
	data.QuerySmsTemplateponse.Code = data.Code
	data.QuerySmsTemplateponse.RequestId = data.RequestId
	data.QuerySmsTemplateponse.Reason = data.Data.ErrDesc
	return data.QuerySmsTemplateponse, nil
}

func (g *UcloudGateway) buildParam(request *http.HttpRequest) {
	request.SetQuery("PublicKey", g.Conf.AppID)
	SignString := utils.GetSignString(request.GetQueryMap(), "") + g.Conf.AppSecret
	t := sha1.New()
	io.WriteString(t, SignString)
	Signature := fmt.Sprintf("%x", t.Sum(nil))
	request.SetQuery("Signature", Signature)
}
