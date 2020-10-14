package models

type UcenterPlatform struct {
	Model
	PlatformName       string `json:"platform_name"`     // 平台名称
	PlatformUsername   string `json:"platform_username"` // 平台账号
	PlatformPassword   string `json:"platform_password"` // 平台密码
	PlatformAppType    int    `json:"platform_app_type"` // 平台类型 0 api 1 web 5 wechat 6 alipay
	PlatformKey        string `json:"platform_key"`      // 平台key 完全标识
	Status             int    `json:"status"`            // 状态 1运行 0暂停维护 -1禁用
	Ak                 string `json:"ak"`
	Sk                 string `json:"sk"`
	PayName            string `json:"pay_name"`
	PayAk              string `json:"pay_ak"`
	PaySk              string `json:"pay_sk"`
	PayNotifyUrl       string `json:"pay_notify_url"` // 支付回调
	PayNotifyFunc      string `json:"pay_notify_func"`
	PlatformSk         string `json:"platform_sk"`
	MessageCallbackUrl string `json:"message_callback_url"` // 消息中心回调地址
	PlatformSecret     string `json:"platform_secret"`      // 平台secret
	UcCallbackUrl      string `json:"uc_callback_url"`      // 用户中心回调地址
}

// 获取平台详情
func PlatformGetInfoOrKey(key string) (*UcenterPlatform, error) {
	var platform UcenterPlatform
	err := db.Where("platform_key = ? AND flag =1", key).First(&platform).Error
	if err != nil {
		return &UcenterPlatform{}, err
	}
	return &platform, nil
}
