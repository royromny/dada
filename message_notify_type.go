package dada

// 消息通知 https://newopen.imdada.cn/#/development/file/messageIndex?_k=xwnkni
type MessageNotify struct {
	MessageType int    `json:"messageType"` // 是	消息类型（1：骑士取消订单推送消息）
	MessageBody string `json:"messageBody"` //	是	消息内容（json字符串）
	CreateTime  int    `json:"createTime"`  //	是	创建时间（unix时间戳）
}

// 骑士取消订单 https://newopen.imdada.cn/#/development/file/transporterCancelOrder?_k=if07yb
type KnightCancel struct {
	OrderId      string `json:"orderId"`      //	是	商家第三方订单号
	DadaOrderId  int64  `json:"dadaOrderId"` //	否	达达订单号
	CancelReason string `json:"cancelReason"` //	是	骑士取消原因
}
