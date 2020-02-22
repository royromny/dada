package dada

// --------------------------------------------------------------------------------
// MessageConfirm 消息确认 https://newopen.imdada.cn/#/development/file/merchantConfirm?_k=6f8hwj
type MessageConfirm struct {
	MessageType int    `json:"messageType"` //  是	消息类型（1：骑士取消订单推送消息）
	MessageBody string `json:"messageBody"` //	是	消息内容（json字符串）
}

type MessageConfirmRsp struct {
	BaseRep
}

// 骑士取消订单 https://newopen.imdada.cn/#/development/file/transporterCancelOrder?_k=if07yb
type KnightCancelConfirm struct {
	OrderId     string `json:"orderId"`     //	是	商家第三方订单号
	DadaOrderId int64  `json:"dadaOrderId"` //	否	达达订单号
	IsConfirm   int    `json:"isConfirm"`   //  是	0:不同意，1:表示同意
}

func (t *MessageConfirm) APIName() string {
	return "/api/message/confirm"
}

func (t *MessageConfirm) Params() map[string]string {
	var m = make(map[string]string)
	return m
}
