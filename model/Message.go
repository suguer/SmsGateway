package model

type Message struct {
	content      string
	templateCode string
	signName     string
	param        map[string]string
	sdkAppId     string
}

func NewMessage(Content string, TemplateCode string, SignName string, Param map[string]string) *Message {
	return &Message{
		content:      Content,
		templateCode: TemplateCode,
		signName:     SignName,
		param:        Param,
	}
}

func (m Message) GetContent() string {
	return m.content
}

func (m Message) GetTemplateCode() string {
	return m.templateCode
}

func (m Message) GetSignName() string {
	return m.signName
}

func (m Message) GetParam() map[string]string {
	return m.param
}

func (m Message) GetSdkAppId() string {
	return m.sdkAppId
}

func (m *Message) SetSdkAppId(val string) {
	m.sdkAppId = val
}
