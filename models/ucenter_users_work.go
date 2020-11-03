package models

import "uc/lib/utils"

type UcenterUsersWork struct {
	Model
	Cuid              int64      `json:"cuid"`
	Name              string     `json:"name"`     // 姓名
	CardId            string     `json:"card_id"`  // 身份证号
	CardTop           string     `json:"card_top"` // 身份证正面
	CardRen           string     `json:"card_ren"` // 手持
	Address           string     `json:"address"`  // 身份证地址
	Nation            string     `json:"nation"`   // 民族
	Gender            string     `json:"gender"`   // 性别
	Province          string     `json:"province"` // 省
	City              string     `json:"city"`     // 市
	Area              string     `json:"area"`     // 区
	IsAuth            int64      `json:"is_auth"`
	WorkType          string     `json:"work_type"`         // 工人类型
	LocationAddress   string     `json:"location_address"`  // 定位地址
	LocationProvince  string     `json:"location_province"` // 定位省
	LocationCity      string     `json:"location_city"`     // 定位市
	LocationArea      string     `json:"location_area"`     // 定位区
	LocationLongitude float64    `json:"location_longitude"`
	LocationLatitude  float64    `json:"location_latitude"`
	OfficeMobile      string     `json:"office_mobile"`
	Miaoshu           string     `json:"miaoshu"`
	GongzuoTime       string     `json:"gongzuo_time"`
	IsGongren         int        `json:"is_gongren"`
	Birth             utils.Time `json:"birth"` // 生日
	IsGongzuo         int        `json:"is_gongzuo"`
	GongrenType       string     `json:"gongren_type"`
	Home              string     `json:"home"`
}

// 新增work
func UsersWorkAdd(d *UcenterUsersWork) error {
	d.Flag = 1
	if err := db.Create(&d).Error; err != nil {
		return err
	}
	return nil
}
