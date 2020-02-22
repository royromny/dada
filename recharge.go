package dada

// Recharge 获取充值链接 https://newopen.imdada.cn/#/development/file/recharge?_k=t5p8m5
func (this *Client) Recharge(param Recharge, sourceId string) (result *RechargeRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}