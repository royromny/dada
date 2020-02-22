package dada

// BalanceQuery 查询账户余额 https://newopen.imdada.cn/#/development/file/balanceQuery?_k=5uinae
func (this *Client) BalanceQuery(param BalanceQuery, sourceId string) (result *BalanceQueryRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}