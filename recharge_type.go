package dada

// --------------------------------------------------------------------------------
// Recharge 获取充值链接 https://newopen.imdada.cn/#/development/file/recharge?_k=t5p8m5
type Recharge struct {
	Amount    float64 `json:"amount"`     //	是	充值金额（单位元，可以精确到分）
	Category  string  `json:"category"`   //	是	生成链接适应场景（category有二种类型值：PC、H5）
	NotifyUrl string  `json:"notify_url"` //	否	支付成功后跳转的页面（支付宝在支付成功后可以跳转到某个指定的页面，微信支付不支持）
}

type RechargeRsp struct {
	BaseRep
	Result string `json:"result"` // 充值的链接
}

func (t *Recharge) APIName() string {
	return "/api/recharge"
}

func (t *Recharge) Params() map[string]string {
	var m = make(map[string]string)
	return m
}
