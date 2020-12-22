package models

import "github.com/jinzhu/gorm"

type Dispatch struct {
	Model
	PartnerId           int          `orm:"partner_id"`    // 合伙人id
	Name                string       `orm:"name"`          // 项目名称
	AdrName             string       `orm:"adr_name"`      // lbs地址名称
	AdrLatitude         float64      `orm:"adr_latitude"`  // 纬度
	AdrLongitude        float64      `orm:"adr_longitude"` // 经度
	Address             string       `orm:"address"`       // 详细地址
	ZxType              string       `orm:"zx_type"`       // 装修类型
	ProductArea         string       `orm:"product_area"`  // 项目面积
	BeginTime           string       `orm:"begin_time"`
	EndTime             string       `orm:"end_time"`
	QuotedType          int          `orm:"quoted_type"` // 报价方式 0自主 1商家
	Price               float64      `orm:"price"`       // 自主报价金额
	IsSjs               int          `orm:"is_sjs"`      // 是否需要设计师
	Describe            string       `orm:"describe"`
	Pics                string       `orm:"pics"`
	Status              int          `orm:"status"` // -1审核没过 0待审核 1过审 5确定接单 9完工
	Cuid                int          `orm:"cuid"`   // 最终接单时 工人id
	User                UcenterUsers `gorm:"ForeignKey:Cuid" json:"user"`
	DispatchNo          string       `orm:"dispatch_no"`
	OrderNo             string       `orm:"order_no"`              // 订单id 预留
	PartnerPrice        float64      `orm:"partner_price"`         // 合伙人实际扣除金额
	IsBack              int          `orm:"is_back"`               // 是否退钱，如果1说明已经把款项退回给合伙人
	UsersServiceId      int          `orm:"users_service_id"`      // 用户服务id
	EndEnrollTime       string       `orm:"end_enroll_time"`       // 报名截止时间
	CheckEndTime        int          `orm:"check_end_time"`        // 工人提交end信息的时间
	IsAdminComment      int          `orm:"is_admin_comment"`      // 是否后台评价
	AdminCommentStar    int          `orm:"admin_comment_star"`    // 评分
	AdminCommentPics    string       `orm:"admin_comment_pics"`    // 评价图片
	AdminCommentContent string       `orm:"admin_comment_content"` // 评价内容
	AdminCommentTime    int          `orm:"admin_comment_time"`
}

// 获取Api列表
func ApiLists(page int, limit int, maps string) ([]*Dispatch, error) {
	offset := (page - 1) * limit
	return DispatchGetLists(offset, limit, maps)
}

// 获取列
func DispatchGetLists(pageNum int, pageSize int, maps interface{}) ([]*Dispatch, error) {
	var d []*Dispatch
	err := db.Model(&Dispatch{}).Preload("User").Where(maps).Offset(pageNum).Limit(pageSize).Order("id desc").Find(&d).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return d, nil
}
func DispatchGetInfo(id int64) (*Dispatch, error) {
	var platform Dispatch
	err := db.Where("id = ? AND flag =1", id).First(&platform).Error
	if err != nil {
		return &Dispatch{}, err
	}
	return &platform, nil
}
