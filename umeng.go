package umeng

import (
	"errors"
)

// Request 基本请求结构
type Request struct {
	Timestamp    string      `json:"timestamp"`
	Type         Mode        `json:"type"`
	DeviceTokens string      `json:"device_tokens"`
	AliasType    string      `json:"alias_type"`
	Alias        string      `json:"alias"`
	FileID       string      `json:"file_id"`
	Description  string      `json:"description"`
	Filter       interface{} `json:"filter"`
	Payload      interface{} `json:"payload"`
	Policy       interface{} `json:"policy"`
	Device       Device      `json:"-"`
}

//Response 基本响应结构
type Response struct {
	Ret  string            `json:"ret"`
	Data map[string]string `json:"data"`
}

// Mode push 方式
type Mode string

const (
	// ModeUnicast 单播
	ModeUnicast Mode = "unicast"
	// ModeListcast 列播
	ModeListcast Mode = "listcast"
	// ModeFilecast 文件播
	ModeFilecast Mode = "filecast"
	// ModeBroadcast 广播
	ModeBroadcast Mode = "broadcast"
	// ModeGroupcast 组播
	ModeGroupcast Mode = "groupcast"
	// ModeCustomizedcast 自定义推送
	ModeCustomizedcast Mode = "customizedcast"
)

func (m Mode) String() string {
	return string(m)
}

// Device 设备类型
type Device int16

const (
	// Unknown 未知设备
	Unknown Device = iota
	// Ios ios 设备
	Ios
	// Android android 设备
	Android
)

// Umeng 主推送对象
type Umeng struct {
	Options
}

// Requester 消息初始化接口
type Requester interface {
	GetRequest() *Request
}

// NewUmeng 实例化友盟推送
func NewUmeng(o *Options) *Umeng {
	return &Umeng{
		Options: *o,
	}
}

// IsIos  是否是 ios 设备
func IsIos(deviceToken string) bool {
	if len(deviceToken) == 64 {
		return true
	}
	return false
}

// IsAndroid 是否是安卓设备
func IsAndroid(deviceToken string) bool {
	if len(deviceToken) == 44 {
		return true
	}
	return false
}

// GetDevice 获取设备类型
func GetDevice(deviceToken string) Device {
	if IsIos(deviceToken) {
		return Ios
	}
	if IsAndroid(deviceToken) {
		return Android
	}
	return Unknown
}

// SetKv 设置自定义字段
func (r *Request) SetKv(k string, v interface{}) error {
	switch r.Device {
	case Ios:
		if val, ok := r.Payload.(IosPayload); ok {
			if val == nil {
				val = make(IosPayload)
			}
			if k == "aps" || k == "d" || k == "p" {
				return errors.New("不允许的键名")
			}
			val[k] = v
			r.Payload = val
		}
		break
	case Android:
		if val, ok := r.Payload.(AndroidPayload); ok {
			if val.Body.Extra == nil {
				val.Body.Extra = make(map[string]interface{})
			}
			val.Body.Extra[k] = v
			val.Body.Custom = val.Body.Extra
			r.Payload = val
		}
		break
	default:
		return errors.New("未知的设备")
	}
	return nil
}
