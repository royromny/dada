package dada

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
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
	fmt.Printf("body, %v\n", string(body))
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
	w.Write([]byte(`{"status":"ok"}`))
}

func VerifySign(notity OrderNotification) (ok bool) {
	var keys []string
	var strParam string
	keys = append(keys, notity.UpdateTime, notity.ClientId, notity.OrderId)
	sort.Strings(keys)
	for i := 0; i < len(keys); i++ {
		strParam += keys[i]
	}
	strParam = MD5(strParam)
	if notity.Signature == strParam {
		return true
	}
	return false
}
