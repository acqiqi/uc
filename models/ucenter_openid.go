package models

type UcenterOpenid struct {
	Model
	Cuid        int64        `json:"cuid"` // ucenter id
	User        UcenterUsers `gorm:"ForeignKey:Cuid" json:"user"`
	PlatformKey string       `json:"platform_key"` // 平台key
	Type        string       `json:"type"`         // 类型 wechat ali app
	Openid      string       `json:"openid"`
}

// 根据openid获取Info
func UcenterOpenidCheckInfo(platform_key, openid string) (*UcenterOpenid, error) {
	var pc UcenterOpenid
	err := db.Where("platform_key = ? AND openid = ? AND flag = 1", platform_key, openid).Preload("User").First(&pc).Error
	if err != nil {
		return &UcenterOpenid{}, err
	}
	return &pc, nil
}

// 新增openid
func UcenterOpenidAdd(d *UcenterOpenid) error {
	d.Flag = 1
	if err := db.Create(&d).Error; err != nil {
		return err
	}
	return nil
}

// 编辑openid
func UcenterOpenidEdit(id int64, data interface{}) error {
	if err := db.Model(&UcenterOpenid{}).Where("id = ? AND flag = 1 ", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

//更新openid
func UcenterOpenidUp(cuid int64, openid, platform_key, openid_type string) bool {
	var o UcenterOpenid
	err := db.Where("platform_key = ? AND cuid = ? AND type = ? AND flag = 1",
		platform_key,
		cuid,
		openid_type).Preload("User").First(&o).Error
	if err != nil {
		//创建Openid
		data := UcenterOpenid{
			Cuid:        cuid,
			PlatformKey: platform_key,
			Type:        openid_type,
			Openid:      openid,
		}
		if err := UcenterOpenidAdd(&data); err != nil {
			return false
		}
	} else {
		o.Openid = openid
		UcenterOpenidEdit(o.Id, o)
	}
	return true
}
