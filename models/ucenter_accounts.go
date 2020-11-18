package models

type UcenterAccounts struct {
	Model
	Cuid        int     `json:"cuid"`         // ucenter uid
	PlatformKey string  `json:"platform_key"` // 牵扯的平台
	Type        int     `json:"type"`         // 0直接收益 1分润收益 2师徒收益 10现金红包收益 11现金卡券 50充值收益 100提现 101购买商品
	Level       int     `json:"level"`        // 收益等级 比如 0一级分享收益 1二级分享收益
	Content     string  `json:"content"`      // 详细内容
	Describe    string  `json:"describe"`     // 描述  主要是显示这里
	ProjectId   int     `json:"project_id"`   // 项目id
	OrderId     int     `json:"order_id"`     // 订单id
	OrderNo     string  `json:"order_no"`     // 订单编号
	Price       float64 `json:"price"`        // 金额
	IsDz        int     `json:"is_dz"`        // 是否到账 1是
	SourceCuid  int     `json:"source_cuid"`  // 来源用户。比如是谁分享产生的给你费用
	ProjectName string  `json:"project_name"` // 项目名称
	Title       string  `json:"title"`        // 标题
	AccountNo   string  `json:"account_no"`   // 订单号
	IsOld       int     `json:"is_old"`       // 是否是老订单
	ProjectType int     `json:"project_type"` // 0上门服务 1招工
}

func UcenterAccountsList(cuid int64) float64 {
	rows := db.Table("vhake_ucenter_accounts").Select("sum(price) as sum_price").
		Where("cuid = ? AND type = 1 AND is_dz = 1 AND flag = 1 AND price < 200", cuid).Row()
	var sum_price float64
	err := rows.Scan(&sum_price)
	if err != nil {
		//return 0, err
		return 0
	}
	return sum_price
}
