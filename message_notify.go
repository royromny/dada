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

// 消息通知 https://newopen.imdada.cn/#/development/file/messageIndex?_k=xwnkni
func (this *Client) GetMessageNotify(req *http.Request) (notity *MessageNotify, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}
	if err = req.ParseForm(); err != nil {
		return nil, err
	}

	notity = &MessageNotify{}
	body, err := ioutil.ReadAll(req.Body)
	if err = json.Unmarshal(body, notity); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}
	return notity, err
}

func (this *Client) AckMessageNotify(w http.ResponseWriter) {
	AckMessageNotify(w)
}

func AckMessageNotify(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
