package router

import (
	"github.com/gin-gonic/gin"
	pv1 "uc/controller/platform/v1"
	v1 "uc/controller/v1"
	"uc/middleware/app_middleware"
	"uc/middleware/jwt"
	"uc/middleware/platform"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(app_middleware.App())
	// 认证类
	//注册AuthRouter
	setAuthRouter(r)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.POST("/get_user_info", v1.GetUserInfo) //获取项目列表
	}
	platform_v1 := r.Group("/platform/v1")
	platform_v1.Use(platform.Platform())
	{
		platform_v1.POST("/mobile_check", pv1.MobileCheck) //手机检测注册
		platform_v1.POST("/mobile_reg", pv1.MobileReg)     //手机注册
	}
	return r
}
