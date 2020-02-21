package dada

// --------------------------------------------------------------------------------
// ShopAdd 新增门店 https://newopen.imdada.cn/#/development/file/shopAdd?_k=lewbit
type ShopAdd struct {
	_ []Shop
}

type Shop struct {
	StationName    string  `json:"station_name"`    //  是	门店名称
	Business       int     `json:"business"`        // 	是	业务类型(食品小吃-1,饮料-2,鲜花-3,文印票务-8,便利店-9,水果生鲜-13,同城电商-19, 医药-20,蛋糕-21,酒品-24,小商品市场-25,服装-26,汽修零配-27,数码-28,小龙虾-29, 其他-5)
	CityName       string  `json:"city_name"`       // 	是	城市名称(如,上海)
	AreaName       string  `json:"area_name"`       // 	是	区域名称(如,浦东新区)
	StationAddress string  `json:"station_address"` // 	是	门店地址
	Lng            float64 `json:"lng"`             // 	是	门店经度
	Lat            float64 `json:"lat"`             // 	是	门店纬度
	ContactName    string  `json:"contact_name"`    // 	是	联系人姓名
	Phone          string  `json:"phone"`           // 	是	联系人电话
	OriginShopId   string  `json:"origin_shop_id"`  // 	否	门店编码,可自定义,但必须唯一;若不填写,则系统自动生成
	IdCard         string  `json:"id_card"`         // 	否	联系人身份证
	Username       string  `json:"username"`        // 	否	达达商家app账号(若不需要登陆app,则不用设置)
	Password       string  `json:"password"`        // 	否	达达商家app密码(若不需要登陆app,则不用设置)
}

type ShopAddRsp struct {
	BaseRep
	Result struct {
		Success     int    `json:"success"`      //	成功导入门店的数量
		SuccessList []Shop `json:"success_list"` //	成功导入的门店
		FailedList  []Shop `json:"failed_list"`  //	导入失败门店编号以及原因
	} `json:"result"`
}

func (t *ShopAdd) APIName() string {
	return "/api/shop/add"
}

func (t *ShopAdd) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// ShopUpdate 编辑门店 https://newopen.imdada.cn/#/development/file/shopUpdate?_k=6396gm
type ShopUpdate struct {
	OriginShopId   string  `json:"origin_shop_id"`  // 	是	门店编码
	NewShopId      string  `json:"new_shop_id"`     // 	否	新的门店编码
	StationName    string  `json:"station_name"`    // 	否	门店名称
	Business       int     `json:"business"`        // 	否	业务类型(食品小吃-1,饮料-2,鲜花-3,文印票务-8,便利店-9,水果生鲜-13,同城电商-19, 医药-20,蛋糕-21,酒品-24,小商品市场-25,服装-26,汽修零配-27,数码-28,小龙虾-29, 其他-5)
	CityName       string  `json:"city_name"`       // 	否	城市名称(如,上海)
	AreaName       string  `json:"area_name"`       // 	否	区域名称(如,浦东新区)
	StationAddress string  `json:"station_address"` // 	否	门店地址
	Lng            float64 `json:"lng"`             // 	否	门店经度
	Lat            float64 `json:"lat"`             // 	否	门店纬度
	ContactName    string  `json:"contact_name"`    // 	否	联系人姓名
	Phone          string  `json:"phone"`           // 	否	联系人电话
	Status         int     `json:"status"`          // 	否	门店状态（1-门店激活，0-门店下线）
}

type ShopUpdateRsp struct {
	BaseRep
}

func (t *ShopUpdate) APIName() string {
	return "/api/shop/update"
}

func (t *ShopUpdate) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// ShopDetail 门店详情 https://newopen.imdada.cn/#/development/file/shopDetail?_k=7k6t2r
type ShopDetail struct {
	OriginShopId string `json:"origin_shop_id"` // 	是	门店编码
}

type ShopDetailRsp struct {
	BaseRep
	Result struct {
		OriginShopId   string  `json:"origin_shop_id"`  // 	是	门店编码
		StationName    string  `json:"station_name"`    // 	否	门店名称
		Business       int     `json:"business"`        // 	否	业务类型(食品小吃-1,饮料-2,鲜花-3,文印票务-8,便利店-9,水果生鲜-13,同城电商-19, 医药-20,蛋糕-21,酒品-24,小商品市场-25,服装-26,汽修零配-27,数码-28,小龙虾-29, 其他-5)
		CityName       string  `json:"city_name"`       // 	否	城市名称(如,上海)
		AreaName       string  `json:"area_name"`       // 	否	区域名称(如,浦东新区)
		StationAddress string  `json:"station_address"` // 	否	门店地址
		Lng            float64 `json:"lng"`             // 	否	门店经度
		Lat            float64 `json:"lat"`             // 	否	门店纬度
		ContactName    string  `json:"contact_name"`    // 	否	联系人姓名
		Phone          string  `json:"phone"`           // 	否	联系人电话
		Status         int     `json:"status"`          // 	否	门店状态（1-门店激活，0-门店下线）
	} `json:"result"`
}

func (t *ShopDetail) APIName() string {
	return "/api/shop/update"
}

func (t *ShopDetail) Params() map[string]string {
	var m = make(map[string]string)
	return m
}
