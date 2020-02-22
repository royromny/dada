package dada

// MessageConfirm 消息确认 https://newopen.imdada.cn/#/development/file/merchantConfirm?_k=6f8hwj
func (this *Client) MessageConfirm(param MessageConfirm, sourceId string) (result *MessageConfirmRsp, err error) {
	err = this.doRequest("POST", &param, sourceId, &result)
	return result, err
}
