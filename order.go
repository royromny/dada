package dada

// AddOrder 新增配送单接口 https://newopen.imdada.cn/#/development/file/add?_k=21pavh
func (this *Client) AddOrder(param AddOrder, sourceId string) (result *AddOrderRsp, err error) {
	err = this.doRequest("POST", param, sourceId, &result)
	return result, err
}

// StatusQuery 订单详情查询 https://newopen.imdada.cn/#/development/file/statusQuery?_k=v8u4ug
func (this *Client) StatusQuery(param StatusQuery, sourceId string) (result *StatusQueryRsp, err error) {
	err = this.doRequest("POST", param, sourceId, &result)
	return result, err
}