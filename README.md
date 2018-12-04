# umeng
友盟 u-push 封装。

本库对友盟 U-Push 推送的消息体进行了 简单归类和再封装。
- 只实现了单播发送，其余的如果有用到仍然需要对本库进行扩展。无非是对 `Request` 类型的各种不同的赋值。具体参看 `unicast.go` 文件的代码实现,还有友盟文档 [U-Push](https://developer.umeng.com/docs/66632/detail/68343#h2--k-18) 。
- 使用示例，发送通知消息
```go
package main
import (
	"git.bitkinetic.com/go/uitls/umeng"
)

var um *umeng.Umeng
var umUnicast *umeng.Unicast

func main() {
	var err error
	um = umeng.NewUmeng(&umeng.Options{
		IosApp: umeng.App{
			AppKey:          "xxxx",
			AppMasterSecret: "xxxx",
		},
		AndroidApp: umeng.App{
			AppKey:          "xxxx",
			AppMasterSecret: "xxxx",
		},
		SSL:    true,
		IsTest: false,
	})
	umUnicast = umeng.NewUnicast()
	umNotice := umUnicast.Notification("deviceToken",  "title", "text", "subTitle", "description")
	err = umNotice.SetKv("k", "value")
	if err != nil {
		// 自定义字段设置错误
	}
	_, err = um.Push(umNotice)
	if err != nil {
		// 发送失败
	}
}
```