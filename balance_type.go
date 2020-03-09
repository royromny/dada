package dada

// --------------------------------------------------------------------------------
// BalanceQuery 查询账户余额 https://newopen.imdada.cn/#/development/file/balanceQuery?_k=5uinae
type BalanceQuery struct {
	Category uint8 `json:"category"` //	否	查询运费账户类型（1：运费账户；2：红包账户，3：所有），默认查询运费账户余额
}

type BalanceQueryRsp struct {
	BaseRep
	Result struct {
		DeliverBalance   float64 `json:"deliverBalance"`   // 交付余额
		RedPacketBalance float64 `json:"redPacketBalance"` // 红包余额
	} `json:"result"`
}

func (t *BalanceQuery) APIName() string {
	return "/api/balance/query"
}

func (t *BalanceQuery) Params() map[string]string {
	var m = make(map[string]string)
	return m
}
