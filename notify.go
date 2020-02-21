package dada

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strconv"
)

// 订单回调
var (
	kSuccess = []byte("success")
)

// 订单回调 https://newopen.imdada.cn/#/development/file/order?_k=md8v7x
func (this *Client) GetOrderNotification(req *http.Request) (notity *OrderNotification, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}
	if err = req.ParseForm(); err != nil {
		return nil, err
	}

	notity = &OrderNotification{}
	body, err := ioutil.ReadAll(req.Body)
	if err = json.Unmarshal(body, notity); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}
	ok := VerifySign(*notity)
	if ok == false {
		return nil, errors.New("VerifySign error")
	}
	return notity, err
}

func (this *Client) AckNotification(w http.ResponseWriter) {
	AckNotification(w)
}

func AckNotification(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write(kSuccess)
}

func VerifySign(notity OrderNotification) (ok bool) {
	// 按key顺序整理出来
	var keys []string
	t := reflect.TypeOf(notity)
	v := reflect.ValueOf(notity)
	for k := 0; k < t.NumField(); k++ {
		keys = append(keys, t.Field(k).Name)
	}
	sort.Strings(keys)
	var strParam string
	for k := 0; k < t.NumField(); k++ {
		if keys[k] != "signature" {
			switch v.FieldByName(keys[k]).Type().Name() {
			case "string":
				strParam += v.FieldByName(keys[k]).String()
			case "int":
				strParam += strconv.Itoa(int(v.FieldByName(keys[k]).Int()))
			}
		}
	}
	strParam = MD5(strParam)
	if notity.Signature == strParam {
		return true
	}
	return false
}
