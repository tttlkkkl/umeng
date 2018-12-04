package umeng

// type IosPayload struct {
// 	Apps IosApps `json:"aps"`
// 	Kv
// }
// IosPayload ios 消息
type IosPayload map[string]interface{}

// IosApps APNs 消息结构
type IosApps struct {
	Alert            IosAlert `json:"alert"`
	Badge            int      `json:"badge"`
	Sound            string   `json:"sound"`
	ContentAvailable int      `json:"content-available"`
	Category         string   `json:"category"`
}

// IosAlert 通知基本信息
type IosAlert struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Body     string `json:"body"`
}

// IosPolicy 可选发送项
type IosPolicy struct {
	StartTime      string `json:"start_time"`
	ExpireTime     string `json:"expire_time"`
	OutBizNo       string `json:"out_biz_no"`
	ApnsCollapseID string `json:"apns_collapse_id"`
}
