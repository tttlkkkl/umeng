package umeng

import (
	"strconv"
	"time"
)

// Message 消息
type Message struct {
	Request
}

// NewMessage 获取消息包装对象
func NewMessage(description string) *Message {
	return &Message{
		Request: Request{
			Timestamp:   strconv.FormatInt(time.Now().Unix(), 10),
			Description: description,
		},
	}
}

// IosMessage 生成 ios 消息对象
func (m *Message) IosMessage(custom interface{}) {
	m.Payload = IosPayload{
		"aps": IosApps{
			ContentAvailable: 1,
		},
		"custom": custom,
	}
}

// AndroidMessage 生成 安卓消息对象
func (m *Message) AndroidMessage(custom interface{}) {
	m.Payload = AndroidPayload{
		DisplayType: "message",
		Body: AndroidBody{
			Custom: custom,
		},
	}
}

// GetRequest 实现Requester接口
func (m *Message) GetRequest() *Request {
	return &m.Request
}
