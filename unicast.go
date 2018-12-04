package umeng

// Unicast 单播发送实现
type Unicast struct {
}

// NewUnicast 获取单播实例
func NewUnicast() *Unicast {
	return &Unicast{}
}

// Message 初始化为一个 message 消息
func (u *Unicast) Message(deviceToken string, custom interface{}, description string) *Message {
	message := NewMessage(description)
	switch GetDevice(deviceToken) {
	case Ios:
		message.IosMessage(custom)
		message.Device = Ios
		break
	case Android:
		message.AndroidMessage(custom)
		message.Device = Android
		break
	default:
		message.Device = Unknown
	}
	message.Type = ModeUnicast
	message.DeviceTokens = deviceToken
	return message
}

// Notification 初始化为一个 notification 通知消息
func (u *Unicast) Notification(deviceToken string, title string, text string, subTitle string, description string) *Notice {
	notice := NewNotice(description)
	switch GetDevice(deviceToken) {
	case Ios:
		notice.IosNotice(title, text, subTitle)
		notice.Device = Ios
		break
	case Android:
		notice.AndroidNotice(title, text, subTitle)
		notice.Device = Android
		break
	default:
		notice.Device = Unknown
	}
	notice.Type = ModeUnicast
	notice.DeviceTokens = deviceToken
	return notice
}
