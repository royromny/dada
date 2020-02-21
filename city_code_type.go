package dada

// --------------------------------------------------------------------------------
// 获取城市信息 https://newopen.imdada.cn/#/development/file/cityList?_k=7xzq6j
type CityCodeList struct{}

type CityCodeListRsp struct {
	BaseRep
	Result []City `json:"result"`
}

type City struct {
	CityName string `json:"cityName"` // 城市名称
	CityCode string `json:"cityCode"` // 城市编码
}

func (t *CityCodeList) APIName() string {
	return "/api/cityCode/list"
}

func (t *CityCodeList) Params() map[string]string {
	var m = make(map[string]string)
	return m
}
