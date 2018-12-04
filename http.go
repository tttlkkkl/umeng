package umeng

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

// 替换原生 json 解码库
var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	//HTTPSendURL http 协议发送地址
	HTTPSendURL = "http://msg.umeng.com/api/send"
	//HTTPSSendURL https 协议发送地址
	HTTPSSendURL = "https://msgapi.umeng.com/api/send"
)

type body struct {
	AppKey         string `json:"appkey"`
	ProductionMode bool   `json:"production_mode"`
	Request
}

// Push 消息推送
func (u *Umeng) Push(r Requester) (*Response, error) {
	var err error
	var appMasterSecret string
	req := r.GetRequest()
	body := &body{
		ProductionMode: !u.IsTest,
		Request:        *req,
	}
	switch req.Device {
	case Ios:
		body.AppKey = u.IosApp.AppKey
		appMasterSecret = u.IosApp.AppMasterSecret
		break
	case Android:
		body.AppKey = u.AndroidApp.AppKey
		appMasterSecret = u.AndroidApp.AppMasterSecret
		break
	default:
		return nil, errors.New(" 未知的设备 ")
	}
	httpBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(httpBody))
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	signStr := "POST" + u.getURL() + string(httpBody) + appMasterSecret
	sign := fmt.Sprintf("%x", md5.Sum([]byte(signStr)))
	bufReader := bytes.NewReader(httpBody)
	resp, err := client.Post(u.getFullURL(sign), "application/json", bufReader)
	if err != nil {
		return nil, err
	}
	defer func() {
		resp.Body.Close()
	}()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var rs Response
	err = json.Unmarshal(content, &rs)
	if err != nil {
		return nil, err
	}
	if rs.Ret == "SUCCESS" {
		return &rs, nil
	}
	var code, msg string
	if e, ok := rs.Data["error_code"]; ok {
		code = e
	}
	if e, ok := rs.Data["error_msg"]; ok {
		msg = e
	}
	return &rs, fmt.Errorf("code:%s , msg: %s", code, msg)
}

func (u *Umeng) getURL() string {
	if u.SSL {
		return HTTPSSendURL
	}
	return HTTPSendURL
}

func (u *Umeng) getFullURL(sign string) string {
	if u.SSL {
		return HTTPSSendURL + "?sign=" + sign
	}
	return HTTPSendURL + "?sign=" + sign
}
