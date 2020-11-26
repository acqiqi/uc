package models

type SmQuotationKey struct {
	Model
	Content  string  `json:"content"`  // 工作内容
	Unit     string  `json:"unit"`     // 计量单位
	Zl       float64 `json:"zl"`       // 主料基础价
	Fl       float64 `json:"fl"`       // 辅料基础价
	Work     float64 `json:"work"`     // 工人基础价
	CatId    int     `json:"cat_id"`   // 大类id
	Describe string  `json:"describe"` // 描述
}

// 获取
func SmQuotationKeyGetInfo(content, describe string) (*SmQuotationKey, error) {
	var sqk SmQuotationKey
	err := db.Where("content = ? AND describe = ? AND flag =1", content, describe).First(&sqk).Error
	if err != nil {
		return &SmQuotationKey{}, err
	}
	return &sqk, nil
}

// 编辑
func SmQuotationKeyEdit(id int64, data interface{}) error {
	if err := db.Model(&SmQuotationKey{}).Where("id = ? AND flag = 1 ", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func SmQuotationKeyAdd(qsk *SmQuotationKey) error {
	qsk.Flag = 1
	if err := db.Create(&qsk).Error; err != nil {
		return err
	}
	return nil
}
