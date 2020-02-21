package dada

// AddOrder 新增配送单接口 https://newopen.imdada.cn/#/development/file/add?_k=21pavh
func (this *Client) AddOrder(param AddOrder, sourceId string) (result *AddOrderRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}

// ReAddOrder 重新发布订单 https://newopen.imdada.cn/#/development/file/reAdd?_k=2jkzyb
func (this *Client) ReAddOrder(param ReAddOrder, sourceId string) (result *ReAddOrderRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}

// QueryDeliverFee 订单预发布-查询订单运费 https://newopen.imdada.cn/#/development/file/readyAdd?_k=0hxqh8
func (this *Client) QueryDeliverFee(param QueryDeliverFee, sourceId string) (result *QueryDeliverFeeRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}

// AddAfterQuery 订单预发布-查询运费后发单 https://newopen.imdada.cn/#/development/file/readyAdd?_k=0hxqh8
func (this *Client) AddAfterQuery(param AddAfterQuery, sourceId string) (result *AddAfterQueryRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}

// AddTip 增加小费 https://newopen.imdada.cn/#/development/file/addTip?_k=l9b2m4
func (this *Client) AddTip(param AddTip, sourceId string) (result *AddTipRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}

// StatusQuery 订单详情查询 https://newopen.imdada.cn/#/development/file/statusQuery?_k=v8u4ug
func (this *Client) StatusQuery(param StatusQuery, sourceId string) (result *StatusQueryRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}

// FormalCancel 取消订单(线上环境) https://newopen.imdada.cn/#/development/file/formalCancel?_k=pxanlf
func (this *Client) FormalCancel(param FormalCancel, sourceId string) (result *FormalCancelRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}

// CancelReasons 获取取消原因 https://newopen.imdada.cn/#/development/file/reasonList?_k=i8gq02
func (this *Client) CancelReasons(param CancelReasons, sourceId string) (result *CancelReasonsRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}

// AppointExist 追加订单 https://newopen.imdada.cn/#/development/file/appointOrder?_k=4ugvgj
func (this *Client) AppointExist(param AppointExist, sourceId string) (result *AppointExistRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}

// AppointCancel 追加订单 https://newopen.imdada.cn/#/development/file/appointOrderCancel?_k=cv6mo9
func (this *Client) AppointCancel(param AppointCancel, sourceId string) (result *AppointCancelRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}

// AppointListTransporter 查询追加配送员 https://newopen.imdada.cn/#/development/file/listTransportersToAppoint?_k=4ecth3
func (this *Client) AppointListTransporter(param AppointListTransporter, sourceId string) (result *AppointListTransporterRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}

// ConfirmGoods 妥投异常之物品返回完成 https://newopen.imdada.cn/#/development/file/abnormalConfirm?_k=fxapv2
func (this *Client) ConfirmGoods(param ConfirmGoods, sourceId string) (result *ConfirmGoodsRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}

