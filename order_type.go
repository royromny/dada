package dada

// --------------------------------------------------------------------------------
// 新增订单 https://newopen.imdada.cn/#/development/file/add?_k=21pavh
type AddOrder struct {
	// 必填
	ShopNo          string  `json:"shop_no"`          // 门店编号，门店创建后可在门店列表和单页查看
	OriginId        string  `json:"origin_id"`        // 第三方订单ID
	CityCode        string  `json:"city_code"`        // 订单所在城市的code（查看各城市对应的code值）
	CargoPrice      float64 `json:"cargo_price"`      // 订单金额
	IsPrepay        uint8   `json:"is_prepay"`        // 是否需要垫付 1:是 0:否 (垫付订单金额，非运费)
	ReceiverName    string  `json:"receiver_name"`    // 收货人姓名
	ReceiverAddress string  `json:"receiver_address"` // 收货人地址
	ReceiverLat     float64 `json:"receiver_lat"`     // 收货人地址维度（高德坐标系）
	ReceiverLng     float64 `json:"receiver_lng"`     // 收货人地址经度（高德坐标系）
	Callback        string  `json:"callback"`         // 回调URL（查看回调说明）
	ReceiverPhone   string  `json:"receiver_phone"`   // 收货人手机号（手机号和座机号必填一项）

	// 非必填
	ReceiverTel    string  `json:"receiver_tel"`     // 收货人座机号（手机号和座机号必填一项）
	Tips           float64 `json:"tips"`             // 小费（单位：元，精确小数点后一位）
	Info           string  `json:"info"`             // 订单备注
	CargoType      uint8   `json:"cargo_type"`       // 订单商品类型：食品小吃-1,饮料-2,鲜花-3,文印票务-8,便利店-9,水果生鲜-13,同城电商-19, 医药-20,蛋糕-21,酒品-24,小商品市场-25,服装-26,汽修零配-27,数码-28,小龙虾-29, 其他-5
	CargoWeight    float64 `json:"cargo_weight"`     // 订单重量（单位：Kg）
	CargoNum       int     `json:"cargo_num"`        // 订单商品数量
	InvoiceTitle   string  `json:"invoice_title"`    // 发票抬头
	OriginMark     string  `json:"origin_mark"`      // 订单来源标示（该字段可以显示在达达app订单详情页面，只支持字母，最大长度为10）
	OriginMarkNo   string  `json:"origin_mark_no"`   // 订单来源编号（该字段可以显示在达达app订单详情页面，支持字母和数字，最大长度为30）
	IsUseInsurance uint8   `json:"is_use_insurance"` // 是否使用保价费（0：不使用保价，1：使用保价； 同时，请确保填写了订单金额（cargo_price））
	// 商品保价费(当商品出现损坏，可获取一定金额的赔付)
	// 保费=配送物品实际价值*费率（5‰），配送物品价值及最高赔付不超过10000元， 最高保费为50元（物品价格最小单位为100元，不足100元部分按100元认定，保价费向上取整数， 如：物品声明价值为201元，保价费为300元*5‰=1.5元，取整数为2元。）
	// 若您选择不保价，若物品出现丢失或损毁，最高可获得平台30元优惠券。 （优惠券直接存入用户账户中）。
	IsFinishCodeNeeded string `json:"is_finish_code_needed"` // 收货码（0：不需要；1：需要。收货码的作用是：骑手必须输入收货码才能完成订单妥投）
	DelayPublishTime   string `json:"delay_publish_time"`    // 预约发单时间（预约时间unix时间戳(10位),精确到分;整10分钟为间隔，并且需要至少提前20分钟预约。）
	IsDirectDelivery   string `json:"is_direct_delivery"`    // 是否选择直拿直送（0：不需要；1：需要。选择直拿直送后，同一时间骑士只能配送此订单至完成，同时，也会相应的增加配送费用）
}

type AddOrderRsp struct {
	BaseRep
	Result struct {
		Distance     float64 `json:"distance"`
		Fee          float64 `json:"fee"`
		DeliverFee   float64 `json:"deliverFee"`
		CouponFee    float64 `json:"couponFee"`
		Tips         float64 `json:"tips"`
		InsuranceFee float64 `json:"insuranceFee"`
	} `json:"result"`
}

func (t *AddOrder) APIName() string {
	return "/api/order/addOrder"
}

func (t *AddOrder) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// ReAddOrder 重新发布订单 https://newopen.imdada.cn/#/development/file/reAdd?_k=2jkzyb
type ReAddOrder struct {
	AddOrder
}

type ReAddOrderRsp struct {
	AddOrderRsp
}

func (t *ReAddOrder) APIName() string {
	return "/api/order/reAddOrder"
}

func (t *ReAddOrder) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// QueryDeliverFee 重新发布订单 https://newopen.imdada.cn/#/development/file/readyAdd?_k=0hxqh8
type QueryDeliverFee struct {
	AddOrder
}

type QueryDeliverFeeRsp struct {
	BaseRep
	Result struct {
		Distance     float64 `json:"distance"`
		DeliveryNo   string  `json:"deliveryNo"` // 平台订单号
		Fee          float64 `json:"fee"`
		DeliverFee   float64 `json:"deliverFee"`
		CouponFee    float64 `json:"couponFee"`
		Tips         float64 `json:"tips"`
		InsuranceFee float64 `json:"insuranceFee"`
	} `json:"result"`
}

func (t *QueryDeliverFee) APIName() string {
	return "/api/order/queryDeliverFee"
}

func (t *QueryDeliverFee) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// AddAfterQuery 订单预发布-查询运费后发单 https://newopen.imdada.cn/#/development/file/readyAdd?_k=0hxqh8
type AddAfterQuery struct {
	DeliveryNo string `json:"deliveryNo"` // 平台订单号
}

type AddAfterQueryRsp struct {
	BaseRep
}

func (t *AddAfterQuery) APIName() string {
	return "/api/order/addAfterQuery"
}

func (t *AddAfterQuery) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// AddTip 增加小费 https://newopen.imdada.cn/#/development/file/addTip?_k=l9b2m4
type AddTip struct {
	DeliveryNo string  `json:"deliveryNo"` // 平台订单号
	OrderId    string  `json:"order_id"`   // 是	第三方订单编号
	Tips       float64 `json:"tips"`       // 是	小费金额(精确到小数点后一位，单位：元)
	CityCode   string  `json:"city_code"`  // 是	订单城市区号
	Info       string  `json:"info"`       // 否	备注(字段最大长度：512)
}

type AddTipRsp struct {
	BaseRep
}

func (t *AddTip) APIName() string {
	return "/api/order/addTip"
}

func (t *AddTip) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// 订单详情查询 https://newopen.imdada.cn/#/development/file/statusQuery?_k=v8u4ug
type StatusQuery struct {
	OrderId string `json:"order_id"` // 第三方订单编号
}

type StatusQueryRsp struct {
	BaseRep
	Result struct {
		OrderId          string  `json:"orderId"`          // 第三方订单编号
		StatusCode       int     `json:"statusCode"`       //	订单状态(待接单＝1 待取货＝2 配送中＝3 已完成＝4 已取消＝5 已过期＝7 指派单=8 妥投异常之物品返回中=9 妥投异常之物品返回完成=10 系统故障订单发布失败=1000 可参考文末的状态说明）
		StatusMsg        string  `json:"statusMsg"`        //	订单状态
		TransporterName  string  `json:"transporterName"`  //	骑手姓名
		TransporterPhone string  `json:"transporterPhone"` //	骑手电话
		TransporterLng   string  `json:"transporterLng"`   //	骑手经度
		TransporterLat   string  `json:"transporterLat"`   //	骑手纬度
		DeliveryFee      float64 `json:"deliveryFee"`      //	配送费,单位为元
		Tips             float64 `json:"tips"`             //	小费,单位为元
		CouponFee        float64 `json:"couponFee"`        //	优惠券费用,单位为元
		InsuranceFee     float64 `json:"insuranceFee"`     //	保价费,单位为元
		ActualFee        float64 `json:"actualFee"`        //	实际支付费用,单位为元
		Distance         float64 `json:"distance"`         //	配送距离,单位为米
		CreateTime       string  `json:"createTime"`       //	发单时间
		AcceptTime       string  `json:"acceptTime"`       //	接单时间,若未接单,则为空
		FetchTime        string  `json:"fetchTime"`        //	取货时间,若未取货,则为空
		FinishTime       string  `json:"finishTime"`       //	送达时间,若未送达,则为空
		CancelTime       string  `json:"cancelTime"`       //	取消时间,若未取消,则为空
		OrderFinishCode  string  `json:"orderFinishCode"`  //	收货码
		DeductFee        float64 `json:"deductFee"`        //	违约金
	} `json:"result"`
}

func (t *StatusQuery) APIName() string {
	return "/api/order/status/query"
}

func (t *StatusQuery) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// FormalCancel 取消订单(线上环境) https://newopen.imdada.cn/#/development/file/formalCancel?_k=pxanlf
type FormalCancel struct {
	OrderId        string `json:"order_id"`         // 是	第三方订单编号
	CancelReasonid int    `json:"cancel_reason_id"` // 是	取消原因ID（取消原因列表）
	CancelReason   string `json:"cancel_reason"`    // 否	取消原因(当取消原因ID为其他时，此字段必填)
}

type FormalCancelRsp struct {
	BaseRep
	Result struct {
		deductfee float64 `json:"deduct_fee"` // 扣除的违约金(单位：元)
	} `json:"result"`
}

func (t *FormalCancel) APIName() string {
	return "/api/order/formalCancel"
}

func (t *FormalCancel) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// CancelReasons 获取取消原因 https://newopen.imdada.cn/#/development/file/reasonList?_k=i8gq02
type CancelReasons struct{}

type CancelReasonsRsp struct {
	BaseRep
	Result []CancelReason `json:"result"`
}

type CancelReason struct {
	Id     int    `json:"id"`     // 理由编号
	Reason string `json:"reason"` // 取消理由
}

func (t *CancelReasons) APIName() string {
	return "/api/order/cancel/reasons"
}

func (t *CancelReasons) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// AppointExist 追加订单 https://newopen.imdada.cn/#/development/file/appointOrder?_k=4ugvgj
type AppointExist struct {
	OrderId       string `json:"order_id"`       // 是	追加的第三方订单ID
	TransporterId int    `json:"transporter_id"` //	是	追加的配送员ID
	ShopNo        string `json:"shop_no"`        // 是	追加订单的门店编码
}

type AppointExistRsp struct {
	BaseRep
}

func (t *AppointExist) APIName() string {
	return "/api/order/appoint/exist"
}

func (t *AppointExist) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// AppointCancel 追加订单 https://newopen.imdada.cn/#/development/file/appointOrderCancel?_k=cv6mo9
type AppointCancel struct {
	OrderId string `json:"order_id"` // 是	追加的第三方订单ID
}

type AppointCancelRsp struct {
	BaseRep
}

func (t *AppointCancel) APIName() string {
	return "/api/order/appoint/cancel"
}

func (t *AppointCancel) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// AppointListTransporter 查询追加配送员 https://newopen.imdada.cn/#/development/file/listTransportersToAppoint?_k=4ecth3
type AppointListTransporter struct {
	ShopNo string `json:"shop_no"` // 是	门店编码
}

type AppointListTransporterRsp struct {
	BaseRep
	Result []Transporter `json:"result"`
}

type Transporter struct {
	Id      int    `json:"id"`      // 配送员id
	Name    string `json:"name"`    // 配送员姓名
	CityIid int    `json:"city_id"` // 配送员城市
}

func (t *AppointListTransporter) APIName() string {
	return "/api/order/appoint/list/transporter"
}

func (t *AppointListTransporter) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// --------------------------------------------------------------------------------
// ConfirmGoods 妥投异常之物品返回完成 https://newopen.imdada.cn/#/development/file/abnormalConfirm?_k=fxapv2
type ConfirmGoods struct {
	OrderId  string `json:"order_id"`  // 是 第三方订单编号
}

type ConfirmGoodsRsp struct {
	BaseRep
}

func (t *ConfirmGoods) APIName() string {
	return "/api/order/confirm/goods"
}

func (t *ConfirmGoods) Params() map[string]string {
	var m = make(map[string]string)
	return m
}
