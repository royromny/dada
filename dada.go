package dada

import (
	"bytes"
	"sort"
	"strconv"
	"strings"
	"time"

	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
)

type Client struct {
	isProduction bool
	apiDomain    string
	Client       *http.Client
	location     *time.Location
	appKey       string
	appSecret    string
}

// New 初始化客户端
// appKey
// appSecret
// isProduction - 是否为生产环境，传 false 的时候为沙箱环境，用于开发测试，正式上线的时候需要改为 true
func New(appKey, appSecret string, isProduction bool) (client *Client, err error) {
	location, _ := time.LoadLocation("Local")
	client = &Client{}
	client.isProduction = isProduction
	client.appKey = appKey
	client.appSecret = appSecret

	if client.isProduction {
		client.apiDomain = kProductionURL
		//client.notifyVerifyDomain = kProductionMAPIURL
	} else {
		client.apiDomain = kSandboxURL
		//client.notifyVerifyDomain = kSandboxURL
	}
	client.Client = http.DefaultClient
	client.location = location

	return client, nil
}

func (this *Client) IsProduction() bool {
	return this.isProduction
}

// 补充处理基本请求参数
func (this *Client) URLValues(param Param, sourceId string) (map[string]string, error) {
	value := make(map[string]string)
	value["app_key"] = this.appKey
	value["timestamp"] = strconv.Itoa(int(time.Now().In(this.location).Unix()))
	value["format"] = kFormat
	value["v"] = kVersion
	value["source_id"] = sourceId // 商户编号（创建商户账号分配的编号） 测试环境默认为：73753

	bytes, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	value["body"] = string(bytes)
	value["signature"] = sign(value, this.appSecret)
	return value, nil
}

func (this *Client) doRequest(method string, param Param, sourceId string, result interface{}) (err error) {
	var buf *bytes.Buffer
	if param != nil {
		p, err := this.URLValues(param, sourceId)
		if err != nil {
			return err
		}
		strP, err := json.Marshal(p)
		buf = bytes.NewBuffer(strP)
	}
	req, err := http.NewRequest(method, this.apiDomain+param.APIName(), buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", kContentType)

	resp, err := this.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("*********dataStr********")
	fmt.Println(string(data))

	err = json.Unmarshal(data, result)
	if err != nil {
		return err
	}

	return err
}

func (this *Client) DoRequest(method string, param Param, sourceId string, result interface{}) (err error) {
	return this.doRequest(method, param, sourceId, result)
}

// 生成签名
func sign(p map[string]string, appSecret string) string {
	// 按key顺序整理出来
	var keys []string
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var strParam string
	for _, k := range keys {
		strParam += k + p[k]
	}
	// 首尾加上app_secret秘钥
	strParam = appSecret + strParam + appSecret
	strParam = strings.ToUpper(MD5(strParam))

	return strParam
}

func MD5(str string) (res string) {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 str
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr) // 输出加密结果
}
