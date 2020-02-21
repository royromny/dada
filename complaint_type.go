package dada

// --------------------------------------------------------------------------------
// Dada 商家投诉达达 https://newopen.imdada.cn/#/development/file/complaintDada?_k=v9oino
type Dada struct {
	OrderId  string `json:"order_id"`  // 是	追加的第三方订单ID
	ReasonId int    `json:"reason_id"` // 是	投诉原因ID（投诉原因列表）
}

type DadaRsp struct {
	BaseRep
}

func (t *Dada) APIName() string {
	return "/api/complaint/dada"
}

func (t *Dada) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// Reasons 获取商家投诉达达原因 https://newopen.imdada.cn/#/development/file/complaintReasons?_k=8jenc0
type Reasons struct {}

type ReasonsRsp struct {
	BaseRep
	Result []Reason `json:"result"`
}

type Reason struct {
	Id     int    `json:"id"`     // 原因编号
	reason string `json:"reason"` // 投诉原因
}

func (t *Reasons) APIName() string {
	return "/api/order/complaint/reasons"
}

func (t *Reasons) Params() map[string]string {
	var m = make(map[string]string)
	return m
}
