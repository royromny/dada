package dada

// MerchantAdd 注册商户 https://newopen.imdada.cn/#/development/file/merchantAdd?_k=rs4djb
func (this *Client) MerchantAdd(param MerchantAdd, sourceId string) (result *MerchantAddRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}
