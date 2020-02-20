package dada

const (
	kSandboxURL    = "http://newopen.qa.imdada.cn"
	kProductionURL = "http://newopen.imdada.cn"
	//kProductionMAPIURL = "https://mapi.alipay.com/gateway.do"

	kFormat      = "json"
	kVersion     = "1.0"
	kContentType = "application/json;charset=utf-8"
)

type Param interface {
	// 用于提供访问的 method
	APIName() string

	// 返回参数列表
	Params() map[string]string
}

type BaseRep struct {
	Status    string `json:"status"`    // 响应状态，成功为"success"，失败为"fail"
	Code      int    `json:"code"`      // 响应返回吗，参考接口返回码
	Msg       string `json:"msg"`       // 响应描述
	ErrorCode int    `json:"errorCode,omitempty"` // 错误编码，与code一致
}
