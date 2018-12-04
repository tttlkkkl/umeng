package umeng

import (
	"strconv"
	"time"
)

// Notice 通知消息
type Notice struct {
	Request
}

// NewNotice 初始化通知消息实例
func NewNotice(description string) *Notice {
	return &Notice{
		Request: Request{
			Timestamp:   strconv.FormatInt(time.Now().Unix(), 10),
			Description: description,
		},
	}
}

// IosNotice 生成 IOS 通知消息对象
func (n *Notice) IosNotice(title string, text string, subTitle string) {
	n.Payload = IosPayload{
		"aps": IosApps{
			Alert: IosAlert{
				Title:    title,
				Subtitle: subTitle,
				Body:     text,
			},
			Sound: "default",
		},
	}
}

// AndroidNotice 生成android 通知消息对象
func (n *Notice) AndroidNotice(title string, text string, subTitle string) {
	n.Payload = AndroidPayload{
		DisplayType: "notification",
		Body: AndroidBody{
			Ticker:    subTitle,
			Title:     title,
			Text:      text,
			AfterOpen: "go_custom",
		},
	}
}

// GetRequest 实现Requester接口
func (n *Notice) GetRequest() *Request {
	return &n.Request
}
