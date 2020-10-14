package go_wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type SmallLoginBody struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    int64  `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

// 微信登录
func SmallLogin(code, appid, secret string) (*SmallLoginBody, error) {
	//获取openid
	requestLine := strings.Join([]string{"https://api.weixin.qq.com/sns/jscode2session",
		"?appid=", appid,
		"&secret=", secret,
		"&js_code=", code,
		"&grant_type=authorization_code"}, "")

	resp, err := http.Get(requestLine)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	logs.Info(err)
	c := string(body)
	log.Println(c)
	if bytes.Contains(body, []byte("openid")) {

	} else {
		return nil, errors.New("get msg fail")
	}

	atr := SmallLoginBody{}

	err = json.Unmarshal(body, &atr)
	if err != nil {
		return nil, err
	} else {
		return &atr, nil
	}
}
