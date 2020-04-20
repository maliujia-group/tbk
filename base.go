package tbk

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"time"
)

type publicConfig struct {
	AppKey       string `json:"appkey"`
	SecretKey    string `json:"secretKey"`
	Sandbox      bool   `json:"sandbox"`
	ApiUri       string `json:"api_uri"`
	ApiBoxUri    string `json:"api_box_uri"`
	Format       string `json:"format"`
	SignMethod   string `json:"signMethod"`
	ApiVersion   string `json:"apiVersion"`
	TargetAPPKey string `json:"target_app_key"`
	PartnerID    string `json:"partner_id"`
	Simplify     bool   `json:"simplify"`
}

type ErrorResponse struct {
	ErrorResponse struct {
		Code      int    `json:"code"`
		Msg       string `json:"msg"`
		SubMsg    string `json:"sub_msg"`
		SubCode   string `json:"sub_code"`
		RequestId string `json:"request_id"`
	} `json:"error_response"`
}

var conf publicConfig

func init() {
	conf = publicConfig{
		ApiUri:     "https://eco.taobao.com/router/rest",
		ApiBoxUri:  "https://gw.api.tbsandbox.com/router/rest",
		Format:     "json",
		SignMethod: "md5",
		ApiVersion: "2.0",
		Sandbox:    false,
		Simplify:   false,
	}
}

type Tbk struct {
	Conf *publicConfig
}

func Client(appKey, secretKey string) *Tbk {
	selfConf := conf
	selfConf.AppKey = appKey
	selfConf.SecretKey = secretKey
	return &Tbk{
		Conf: &selfConf,
	}
}

func (t *Tbk) SetSandBox(sb bool) *Tbk {
	t.Conf.Sandbox = sb
	return t
}

func (t *Tbk) SetFormat(format string) *Tbk {
	t.Conf.Format = format
	return t
}

// 获取系统公共请求参数
func (t *Tbk) systemParams(method string) map[string]string {
	return map[string]string{
		"method":      method,
		"app_key":     t.Conf.AppKey,
		"v":           t.Conf.ApiVersion,
		"format":      t.Conf.Format,
		"sign_method": t.Conf.SignMethod,
		"timestamp":   time.Now().Format("2006-01-02 15:04:05"),
	}
}

// 获取请求URI
func (t *Tbk) getUri(params map[string]string) string {
	uri := fmt.Sprintf("%s?", t.Conf.ApiUri)
	for key, val := range params {
		uri += fmt.Sprintf("%s=%s&", key, url.QueryEscape(val))
	}
	return uri[:len(uri)-1]
}

func (t *Tbk) ParseBody(body io.Reader, response interface{}) ([]byte, error) {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	errRes := ErrorResponse{}
	if err := json.Unmarshal(data, &errRes); err == nil && errRes.ErrorResponse.Code != 0 {
		return nil, errors.New(fmt.Sprintf("%s 错误码：%s", errRes.ErrorResponse.SubMsg, errRes.ErrorResponse.SubCode))
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}
	return data, nil

}

func (t *Tbk) httpPost(method string, data map[string]string, response interface{}) ([]byte, error) {
	sysParams := t.systemParams(method)
	allParams := sysParams
	urlVals := url.Values{}
	for dataKey, dataVal := range data {
		allParams[dataKey] = dataVal
		urlVals.Add(dataKey, dataVal)
	}
	sysParams["sign"] = t.generateSign(allParams, t.Conf.SecretKey)
	requireUri := t.getUri(sysParams)

	resp, err := http.PostForm(requireUri, urlVals)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return t.ParseBody(resp.Body, &response)
}

func (t *Tbk) httpPostJson(method string, data interface{}, response interface{}) ([]byte, error) {
	sysParams := t.systemParams(method)
	sysParams["sign"] = t.generateSign(sysParams, t.Conf.SecretKey)
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", t.getUri(sysParams), bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return t.ParseBody(resp.Body, &response)
}

// 签名算法
func (t *Tbk) generateSign(arr map[string]string, secretKey string) string {
	// 排序
	var keys []string
	for k, _ := range arr {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	stringToBeSigned := secretKey
	for _, key := range keys {
		stringToBeSigned += fmt.Sprintf("%s%s", key, arr[key])
	}
	stringToBeSigned += secretKey
	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(stringToBeSigned))))
}

// 将结构体转换成 map[string]string
func (t *Tbk) Struct2MapString(obj interface{}) map[string]string {
	tp := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]string)
	for i := 0; i < tp.NumField(); i++ {
		key := tp.Field(i).Tag.Get("json")
		if key == "" {
			key = tp.Field(i).Name
		}
		data[key] = v.Field(i).String()
	}
	return data
}
