package gateway

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/suguer/SmsGateway/model"
	"github.com/suguer/SmsGateway/private/http"
	"github.com/suguer/SmsGateway/private/utils"
)

type TencentGateway struct {
	Gateway
}

type TencentSendStatusSetResponse struct {
	SerialNo       string
	PhoneNumber    string
	SessionContext string
	Code           string
	Message        string
	IsoCode        string
	Fee            int
}

type TencentSendSMSMessageResponse struct {
	model.SendSMSMessageResponse
	Response struct {
		RequestId string
		Error     struct {
			Code    string
			Message string
		}
		SendStatusSet []TencentSendStatusSetResponse
	}
}

type TencentCreateSmsTemplateResponse struct {
	model.CreateSmsTemplateResponse
	Response struct {
		RequestId string
		Error     struct {
			Code    string
			Message string
		}
	}
}

func (g *TencentGateway) GetName() string {
	return "tencent"
}

func (g *TencentGateway) Init(c *model.Config) {
	g.Gateway.Init(c)
	g.ApiUrl = "https://sms.tencentcloudapi.com"

}

func (g *TencentGateway) SendMessage(mobile *model.Phone, message *model.Message) (model.SendSMSMessageResponse, error) {
	var data TencentSendSMSMessageResponse
	request := http.NewHttpRequest()
	request.SetMethod("POST")
	request.SetHeader("X-TC-Action", "SendSms")
	TemplateParamSet := make([]string, len(message.GetParam()))
	for k, v := range message.GetParam() {
		index, _ := strconv.Atoi(k)
		TemplateParamSet[index] = v
	}
	args := map[string]any{
		"PhoneNumberSet":   []string{mobile.GetNumber()},
		"SmsSdkAppId":      message.GetSdkAppId(),
		"TemplateId":       message.GetTemplateCode(),
		"SignName":         message.GetSignName(),
		"TemplateParamSet": TemplateParamSet,
	}

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
	data.SendSMSMessageResponse.RequestId = data.Response.RequestId
	if len(data.Response.SendStatusSet) > 0 {
		data.SendSMSMessageResponse.BizId = data.Response.SendStatusSet[0].SerialNo
		data.SendSMSMessageResponse.Message = data.Response.SendStatusSet[0].Message
		data.SendSMSMessageResponse.Code = model.Code{Val: data.Response.SendStatusSet[0].Code}
	} else if data.Response.Error.Message != "" {
		data.SendSMSMessageResponse.Message = data.Response.Error.Message
		data.SendSMSMessageResponse.Code = model.Code{Val: data.Response.Error.Code}
	}
	if len(data.Response.Error.Message) == 0 {
		data.SendSMSMessageResponse.Code.Val = "OK"
	}
	return data.SendSMSMessageResponse, nil
}

func (g *TencentGateway) CreateSmsTemplate(template *model.Template) (model.CreateSmsTemplateResponse, error) {
	var data TencentCreateSmsTemplateResponse
	request := http.NewHttpRequest()
	request.SetMethod("POST")
	request.SetHeader("X-TC-Action", "AddSmsTemplate")
	SmsType := 0
	if template.GetType2Int() != 1 && template.GetType2Int() != 2 {
		return data.CreateSmsTemplateResponse, errors.New("type.unsupport")
	} else if template.GetType2Int() == 1 {
		SmsType = 0
	} else {
		SmsType = 1
	}
	args := map[string]any{
		"TemplateName":    template.GetName(),
		"TemplateContent": template.GetContent(),
		"SmsType":         SmsType,
		"International":   template.GetInternational(),
		"Remark":          template.GetRemark(),
	}
	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return data.CreateSmsTemplateResponse, err
	}
	request.SetRequestBody(jsonBytes)
	g.buildParam(request)
	response, err := g.send(request)
	err = model.NewCommonResponse(&data, response.GetBody())
	if err != nil {
		return data.CreateSmsTemplateResponse, err
	}
	if len(data.Response.Error.Message) != 0 {
		data.Code = model.Code{Val: data.Response.Error.Code}
		data.Message = data.Response.Error.Message
		return data.CreateSmsTemplateResponse, errors.New(data.Message)
	}
	return data.CreateSmsTemplateResponse, nil
}

func (g *TencentGateway) buildParam(request *http.HttpRequest) {

	service := "sms"
	version := "2021-01-11"
	action := request.GetHeader("X-TC-Action")

	format, _ := url.Parse(g.ApiUrl)
	Host := format.Host
	region := "ap-guangzhou"
	Timestamp := time.Now().Unix()
	RequestTimestamp := time.Unix(Timestamp, 0).UTC().Format("2006-01-02")
	CanonicalURI := "/"
	CanonicalQueryString := ""
	CanonicalHeaders := fmt.Sprintf("content-type:%s\nhost:%s\n", http.MimeJSON, Host)
	SignedHeaders := "content-type;host"
	HashedRequestPayload := utils.SHA256Hex(string(request.GetRequestBody()))
	CanonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		request.GetMethod(),
		CanonicalURI,
		CanonicalQueryString,
		CanonicalHeaders,
		SignedHeaders,
		HashedRequestPayload)
	Algorithm := "TC3-HMAC-SHA256"
	CredentialScope := fmt.Sprintf("%s/sms/tc3_request", RequestTimestamp)
	HashedCanonicalRequest := utils.SHA256Hex(CanonicalRequest)
	StringToSign := fmt.Sprintf("%s\n%d\n%s\n%s", Algorithm, Timestamp, CredentialScope, HashedCanonicalRequest)
	secretDate := hmacsha256(RequestTimestamp, "TC3"+g.Conf.AppSecret)
	secretService := hmacsha256(service, secretDate)
	secretSigning := hmacsha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(hmacsha256(StringToSign, secretSigning)))
	authorization := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		Algorithm,
		g.Conf.AppID,
		CredentialScope,
		SignedHeaders,
		signature)
	s := strconv.FormatInt(Timestamp, 10)
	request.SetHeader(http.HeaderNameContentType, http.MimeJSON)
	request.SetHeader("Host", Host)
	request.SetHeader("X-TC-Action", action)
	request.SetHeader("X-TC-Timestamp", s)
	request.SetHeader("X-TC-Version", version)
	request.SetHeader("X-TC-Region", region)
	request.SetHeader("Authorization", authorization)

}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}
