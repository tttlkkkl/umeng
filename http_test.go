package umeng

import (
	"reflect"
	"testing"
)

func TestUmeng_Push(t *testing.T) {
	type args struct {
		r Requester
	}
	iosToken := "562007df7aaae45cdbe4e7ab3d5231be48086ecdd5fabfb6c7764b2d42e87ccb"
	androidToken := "Al78vkl6LWLwQse9vuKSXqeKDHdjkQM_X8hUqEM0UNr5"
	iosApp := App{
		AppKey:          "yyyyy",
		AppMasterSecret: "yyyyyyyy",
	}
	androidApp := App{
		AppKey:          "xxx",
		AppMasterSecret: "xxxxxx",
	}
	umeng := NewUmeng(&Options{
		IosApp:     iosApp,
		AndroidApp: androidApp,
		SSL:        false,
		IsTest:     false,
	})
	umengSSL := NewUmeng(&Options{
		IosApp:     iosApp,
		AndroidApp: androidApp,
		SSL:        true,
		IsTest:     false,
	})
	unicast := NewUnicast()
	argsAndroidMessage := unicast.Message(androidToken, map[string]string{"k": "v"}, "测试发送")
	argsIosMessage := unicast.Message(iosToken, map[string]string{"k": "v"}, "测试发送")
	argsAndroidNotice := unicast.Notification(androidToken, "消息标题-a", "这是消息内容-a", "消息副标题-a", "测试发送")
	argsIOSNotice := unicast.Notification(iosToken, "消息标题", "这是消息内容", "消息副标题", "测试发送")
	argsIosMessage.SetKv("kk", "vv")
	argsIOSNotice.SetKv("kk", "vv")
	tests := []struct {
		name    string
		u       *Umeng
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "android-Message",
			u:    umeng,
			args: args{argsAndroidMessage},
		},
		{
			name: "ios-Message",
			u:    umeng,
			args: args{argsIosMessage},
		},
		{
			name: "android-Notice",
			u:    umeng,
			args: args{argsAndroidNotice},
		},
		{
			name: "ios-Notice",
			u:    umeng,
			args: args{argsIOSNotice},
		},
		{
			name: "android-MessageSSL",
			u:    umengSSL,
			args: args{argsAndroidMessage},
		},
		{
			name: "ios-MessageSSL",
			u:    umengSSL,
			args: args{argsIosMessage},
		},
		{
			name: "android-NoticeSSL",
			u:    umengSSL,
			args: args{argsAndroidNotice},
		},
		{
			name: "ios-NoticeSSL",
			u:    umengSSL,
			args: args{argsIOSNotice},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.Push(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Umeng.Push() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Umeng.Push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUmeng_getURL(t *testing.T) {
	tests := []struct {
		name string
		u    *Umeng
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.getURL(); got != tt.want {
				t.Errorf("Umeng.getURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUmeng_getFullURL(t *testing.T) {
	type args struct {
		sign string
	}
	tests := []struct {
		name string
		u    *Umeng
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.getFullURL(tt.args.sign); got != tt.want {
				t.Errorf("Umeng.getFullURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
