package models

import "github.com/jinzhu/gorm"

type MarketingPacket struct {
	Model
	Cuid        int          `json:"cuid"`
	User        UcenterUsers `gorm:"ForeignKey:Cuid" json:"user"`
	Price       float64      `json:"price"`
	Title       string       `json:"title"`
	Describe    string       `json:"describe"` // 描述
	Status      int          `json:"status"`   // 0正常 1使用 -1过期或者禁用
	Type        string       `json:"type"`     // 来源渠道
	PlatformKey string       `json:"platform_key"`
	PacketNo    string       `json:"packet_no"`
	SCuid       int          `json:"s_cuid"` // 来源用户
	SUser       UcenterUsers `gorm:"ForeignKey:SCuid" json:"s_user"`
	EndTime     int          `json:"end_time"`
}

func PacketGetAll() ([]*MarketingPacket, error) {
	var mb []*MarketingPacket
	err := db.Model(&MarketingPacket{}).Preload("User").Preload("SUser").Order("id asc").Find(&mb).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return mb, nil
}
