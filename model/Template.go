package model

import (
	"strconv"
)

type Template struct {
	name    string
	content string
	// 0：验证码 1：短信通知 2：推广短信。
	templateType string
	remark       string
	// 是否国际/港澳台短信： 0：表示国内短信。 1：表示国际/港澳台短信。
	international string
}

func NewTemplate(name string, content string, templateType string, remark string, international string) *Template {
	return &Template{
		name:          name,
		content:       content,
		templateType:  templateType,
		remark:        remark,
		international: international,
	}
}

func (m Template) GetName() string {
	return m.name
}

func (m Template) GetContent() string {
	return m.content
}

func (m Template) GetType() string {
	return m.templateType
}

func (m Template) GetType2Int() int {
	templateType, _ := strconv.Atoi(m.templateType)
	return templateType
}

func (m Template) GetRemark() string {
	return m.remark
}

func (m Template) GetInternational() string {
	return m.international
}

func (m Template) GetInternational2Int() int {
	international, _ := strconv.Atoi(m.international)
	return international
}
