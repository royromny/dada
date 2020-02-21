package dada

// Dada 商家投诉达达 https://newopen.imdada.cn/#/development/file/complaintDada?_k=v9oino
func (this *Client) Dada(param Dada, sourceId string) (result *DadaRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}

// Reasons 获取商家投诉达达原因 https://newopen.imdada.cn/#/development/file/complaintReasons?_k=8jenc0
func (this *Client) Reasons(param Reasons, sourceId string) (result *ReasonsRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}
