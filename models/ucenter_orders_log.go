package models

import "github.com/jinzhu/gorm"

type UcenterOrdersLog struct {
	Model
	Cuid        int64   `json:"cuid"`
	PlatformKey string  `json:"platform_key"`
	OrderId     int64   `json:"order_id"`
	OrderNo     string  `json:"order_no"`
	Type        int     `json:"type"`         // 0平台收益 1商家收益(工人) 2分润收入 3师徒收入 10合伙人收入
	ProjectType int     `json:"project_type"` // 项目类型 0上门服务
	ProjectId   int64   `json:"project_id"`   // 对应项目id us_id
	Describe    string  `json:"describe"`     // 描述
	Price       float64 `json:"price"`        // 对应金额
}

// 获取列
func OrdersLogGetGaiList() ([]*UcenterOrdersLog, error) {
	var d []*UcenterOrdersLog
	err := db.Model(&UcenterOrdersLog{}).Where("`describe` like ? AND project_id > 0 AND project_type = 1", "%二级师徒收益%").Order("id desc").Find(&d).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return d, nil
}

func OrdersLogEdit(id int64, data interface{}) error {
	if err := db.Model(&UcenterOrdersLog{}).Where("id = ? AND flag = 1 ", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
