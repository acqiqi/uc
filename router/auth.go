package router

import (
	"github.com/gin-gonic/gin"
	"uc/controller/auth"
	"uc/controller/platform"
)

func setAuthRouter(r *gin.Engine) {
	// 平台接口集
	r.POST("/platform_auth/get_access_token", platform.GetAccessToken)

	// 实现登录
	r.POST("/wechat_auth/get_access_token", auth.GetMPOpenId)
	r.POST("/wechat_auth/small_quck_login", auth.SmallQuckLogin)                   //小程序快速登录
	r.POST("/wechat_auth/small_quck_mobile_reg", auth.SmallQuckMobileReg)          //小程序快速注册
	r.POST("/wechat_auth/small_bind_user_info_login", auth.SmallBindUserInfoLogin) //小程序快速绑定用户信息

	// Mobile
	r.POST("/mobile_auth/auto_reg", auth.MobileAutoReg)      //小程序快速绑定用户信息
	r.POST("/mobile_auth/get_reg_sms", auth.MobileGetRegSMS) //小程序快速绑定用户信息

}
