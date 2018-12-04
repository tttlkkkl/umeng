package umeng

// Options 可选项定义
type Options struct {
	IosApp     App
	AndroidApp App
	// 是否启用 https 连接
	SSL bool
	// IsTest 是否是测试发送
	IsTest bool
}

// App 友盟应用配置
type App struct {
	AppKey          string
	AppMasterSecret string
}
