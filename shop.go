package dada

// ShopAdd 新增门店 https://newopen.imdada.cn/#/development/file/shopAdd?_k=lewbit
func (this *Client) ShopAdd(param ShopAdd, sourceId string) (result *ShopAddRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}

// ShopUpdate 编辑门店 https://newopen.imdada.cn/#/development/file/shopUpdate?_k=6396gm
func (this *Client) ShopUpdate(param ShopUpdate, sourceId string) (result *ShopUpdateRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}

// ShopDetail 门店详情 https://newopen.imdada.cn/#/development/file/shopDetail?_k=7k6t2r
func (this *Client) ShopDetail(param ShopDetail, sourceId string) (result *ShopDetailRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}
