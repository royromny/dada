package dada

// --------------------------------------------------------------------------------
// MerchantAdd 注册商户 https://newopen.imdada.cn/#/development/file/merchantAdd?_k=rs4djb
type MerchantAdd struct {
	Mobile            string `json:"mobile"`             // 是	注册商户手机号,用于登陆商户后台
	CityName          string `json:"city_name"`          // 是	商户城市名称(如,上海)
	EnterpriseName    string `json:"enterprise_name"`    // 是	企业全称
	EnterpriseAddress string `json:"enterprise_address"` // 是	企业地址
	ContactName       string `json:"contact_name"`       // 是	联系人姓名
	ContactPhone      string `json:"contact_phone"`      // 是	联系人电话
	Email             string `json:"email"`              // 是	邮箱地址
}

type MerchantAddRsp struct {
	BaseRep
}

func (t *MerchantAdd) APIName() string {
	return "/merchantApi/merchant/add"
}

func (t *MerchantAdd) Params() map[string]string {
	var m = make(map[string]string)
	return m
}
