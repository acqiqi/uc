package models

import "github.com/jinzhu/gorm"

type SmQuotationTplLink struct {
	Model
	TplId     int    `json:"tpl_id"`    // 模板id
	CatId     int    `json:"cat_id"`    // 模板分类id
	CatName   string `json:"cat_name"`  // 模板分类名称
	Describe  string `json:"describe"`  // 备注
	TotalNum  int    `json:"total_num"` // 项目数
	TableJson string `json:"table_json"`
	LinkNo    string `json:"link_no"` // 链接编号
}

func QuotationTplLinkAllList() ([]*SmQuotationTplLink, error) {
	var mb []*SmQuotationTplLink
	err := db.Model(&SmQuotationTplLink{}).Where("flag = 1").Order("id asc").Find(&mb).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return mb, nil
}
