package model

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
)

type Type int

const (
	Int Type = iota
	String
)

type Code struct {
	Val string
}

// 实现 json.Unmarshaller 接口
func (code *Code) UnmarshalJSON(value []byte) error {
	if value[0] == '"' {
		return json.Unmarshal(value, &code.Val)
	}
	var temp int
	err := json.Unmarshal(value, &temp)
	if err != nil {
		return err
	}
	code.Val = strconv.Itoa(temp)
	return nil
}
func (code *Code) GetCode() string {
	return code.Val
}

type CommonResponse struct {
	Message    string
	Code       Code
	RequestId  string
	RawContent string
}

func (m *CommonResponse) SetRawContent(content string) {
	m.RawContent = content
}

func NewCommonResponse(data interface{}, body []byte) error {
	typeOfData := reflect.ValueOf(data)
	err := json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	method := typeOfData.MethodByName("SetRawContent")
	if method.IsValid() {
		in := make([]reflect.Value, 1)
		in[0] = reflect.ValueOf(string(body))
		method.Call(in)
	} else {
		return errors.New("RawContent.empty")
	}
	return nil
}

type SendSMSMessageResponse struct {
	CommonResponse
	BizId string
}

type CreateSmsTemplateResponse struct {
	CommonResponse
	TemplateCode string
}

func (r *CreateSmsTemplateResponse) transfer() {

}

type QuerySmsTemplateponse struct {
	CommonResponse
	// 0：审核中。 1：审核通过。 2：审核失败，请在返回参数Reason中查看审核失败原因。 10：取消审核。
	// 3 待审核,4 被禁用
	TemplateStatus int
	Reason         string
}
