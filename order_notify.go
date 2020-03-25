package dada

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
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
	if errNotity := json.Unmarshal(body, notity); errNotity != nil {
		fmt.Printf("Unmarshal errNotity, %v\n", errNotity)
		// 对付变态的updateTime不同类型问题
		notity2 := &OrderNotification2{}
		if errNotity2 := json.Unmarshal(body, notity2); errNotity2 != nil {
			fmt.Printf("Unmarshal errNotity2, %v\n", errNotity2)
			return nil, errNotity2
		}
		notity.ClientId = notity2.ClientId
		notity.OrderId = notity2.OrderId
		notity.OrderStatus = notity2.OrderStatus
		notity.CancelReason = notity2.CancelReason
		notity.CancelFrom = notity2.CancelFrom
		notity.Signature = notity2.Signature
		notity.DmId = notity2.DmId
		notity.DmName = notity2.DmName
		notity.DmMobile = notity2.DmMobile
		notity.UpdateTime = strconv.Itoa(notity2.UpdateTime)
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
