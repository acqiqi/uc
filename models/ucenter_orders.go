package models

type UcenterOrders struct {
	Model
	OrderNo        string  `json:"order_no"`       // 订单编号
	PlatformKey    string  `json:"platform_key"`   // 平台key
	Cuid           int     `json:"cuid"`           // ucenter id
	CouponKey      string  `json:"coupon_key"`     // 优惠券key
	CouponId       int     `json:"coupon_id"`      // 优惠券id
	CouponPrice    float64 `json:"coupon_price"`   // 优惠券抵扣金额
	CostPrice      float64 `json:"cost_price"`     // 原价
	UnitPrice      float64 `json:"unit_price"`     // 单价
	Price          float64 `json:"price"`          // 现价 支付总价
	GoodsNum       float64 `json:"goods_num"`      // 商品总数量
	PayType        int     `json:"pay_type"`       // 支付类型 0线上 1线下
	PayPlatform    string  `json:"pay_platform"`   // 支付平台 wechat alipay wallet ...
	Status         int     `json:"status"`         // 状态 0下单 1线下付款提交 2审核通过 7已接单 8已发货 9已结算或者已收货
	IsPay          int     `json:"is_pay"`         // 是否支付 0否 1线下提交 2已支付
	IsUComment     int     `json:"is_u_comment"`   // 是否用户评论 0否 1是
	IsMComment     int     `json:"is_m_comment"`   // 是否商家评论 0否1是
	PayTime        int     `json:"pay_time"`       // 支付时间
	ServicePrice   float64 `json:"service_price"`  // 服务费
	SharePrice     float64 `json:"share_price"`    // 分享出去多少钱
	PlatformPrice  float64 `json:"platform_price"` // 平台受益
	ShareLv1       float64 `json:"share_lv1"`      // 一级分享
	ShareLv2       float64 `json:"share_lv2"`      // 二级分享
	OptLv1         float64 `json:"opt_lv1"`        // 一级其他分享
	OptLv2         float64 `json:"opt_lv2"`        // 二级其他分享
	ShareLv1Cuid   int     `json:"share_lv1_cuid"` // 一级分享用户id
	ShareLv2Cuid   int     `json:"share_lv2_cuid"`
	OptLv1Cuid     int     `json:"opt_lv1_cuid"`
	OptLv2Cuid     int     `json:"opt_lv2_cuid"`
	IsAccounts     int     `json:"is_accounts"` // 是否结算
	OrderType      int     `json:"order_type"`  // 订单类型0常规购买订单  1团购 10VIP 20快捷支付
	PayNo          string  `json:"pay_no"`      // 线上付款订单
	Describe       string  `json:"describe"`    // 描述
	ProjectId      int     `json:"project_id"`
	VipPriceD      float64 `json:"vip_price_d"` // vip折扣多少钱
	Pics           string  `json:"pics"`
	PartnerPrice   float64 `json:"partner_price"`    // 合伙人金额
	PartnerId      int     `json:"partner_id"`       // 合伙人id
	TotalUsePrice  float64 `json:"total_use_price"`  // 全部占用金额
	UserPayDesc    string  `json:"user_pay_desc"`    // 用户线下支付备注
	MainOrderId    int     `json:"main_order_id"`    // 关联订单id
	TotalLinkPrice float64 `json:"total_link_price"` // 所有关联订单金额
	UseShareLv2    int     `json:"use_share_lv2"`    // 强制二级分享
	UseShareLv1    int     `json:"use_share_lv1"`    // 强制一级分享
}

func UcenterOrdersGetInfo(id int64) (*UcenterOrders, error) {
	var platform UcenterOrders
	err := db.Where("id = ? AND flag =1", id).First(&platform).Error
	if err != nil {
		return &UcenterOrders{}, err
	}
	return &platform, nil
}
