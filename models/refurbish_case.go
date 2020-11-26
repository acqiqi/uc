package models

type RefurbishCase struct {
	Model
	Title               string `json:"title"`         // 标题
	Logo                string `json:"logo"`          // 首页图
	ShowNum             int    `json:"show_num"`      // 访问数量
	RoomPattern         string `json:"room_pattern"`  // 户型
	CoverImgUrl         string `json:"cover_img_url"` // 详情页头图
	Description         string `json:"description"`   // 描述
	Status              int    `json:"status"`        // 状态 1正常 0隐藏
	ToPath              string `json:"to_path"`       // 转跳页面
	CreatedUser         string `json:"created_user"`  // 发布用户
	UpdateUser          string `json:"update_user"`
	Village             string `json:"village"` // 地址
	AvatarUrl           string `json:"avatar_url"`
	Nickname            string `json:"nickname"`
	Content             string `json:"content"`
	StyleName           string `json:"style_name"`            // 风格
	Area                int    `json:"area"`                  // 面积
	ApartmentLayoutName string `json:"apartment_layout_name"` // 户型
}

func RefurbishCaseAdd(qsk *RefurbishCase) error {
	qsk.Flag = 1
	if err := db.Create(&qsk).Error; err != nil {
		return err
	}
	return nil
}
