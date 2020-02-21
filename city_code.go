package dada

// CityCodeList 获取城市信息 https://newopen.imdada.cn/#/development/file/cityList?_k=7xzq6j
func (this *Client) CityCodeList(param CityCodeList, sourceId string) (result *CityCodeListRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}