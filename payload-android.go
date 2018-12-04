package umeng

// AndroidPayload 安卓消息体
type AndroidPayload struct {
	DisplayType string      `json:"display_type"`
	Body        AndroidBody `json:"body"`
	Mipush      bool        `json:"mipush"`
}

// AndroidBody android 消息体定义
type AndroidBody struct {
	Ticker      string                 `json:"ticker"`
	Title       string                 `json:"title"`
	Text        string                 `json:"text"`
	Icon        string                 `json:"icon"`
	LargeIcon   string                 `json:"largeIcon"`
	Img         string                 `json:"img"`
	Sound       string                 `json:"sound"`
	BuilderID   int                    `json:"builder_id"`
	PlayVibrate bool                   `json:"play_vibrate"`
	PlayLights  bool                   `json:"play_lights"`
	PlaySound   bool                   `json:"play_sound"`
	AfterOpen   string                 `json:"after_open"`
	URL         string                 `json:"url"`
	Activity    string                 `json:"activity"`
	Custom      interface{}            `json:"custom"`
	Extra       map[string]interface{} `json:"extra"`
}

// AndroidPolicy 可选发送项
type AndroidPolicy struct {
	StartTime  string `json:"start_time"`
	ExpireTime string `json:"expire_time"`
	MaxSendNum int    `json:"max_send_num"`
	OutBizNo   string `json:"out_biz_no"`
}
