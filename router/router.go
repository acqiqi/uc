package router

import (
	"github.com/gin-gonic/gin"
	v1 "uc/controller/v1"
	"uc/middleware/app_middleware"
	"uc/middleware/jwt"
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
	return r
}
