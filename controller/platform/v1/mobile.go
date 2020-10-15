package v1

import (
	"github.com/gin-gonic/gin"
	"uc/lib/e"
	"uc/models"
)

// 检测用户是否有注册
func MobileCheck(c *gin.Context) {
	data := struct {
		Mobile string `json:"mobile"`
	}{}
	if err := c.BindJSON(&data); err != nil {
		e.ApiErr(c, err.Error())
		return
	}

	if user, err := models.UsersCheckMobile(data.Mobile); err != nil {
		e.ApiErr(c, "查无用户")
		return
	} else {
		e.ApiOk(c, "获取成功", struct {
			UserInfo models.UcenterUsers `json:"user_info"`
		}{
			UserInfo: *user,
		})
	}
}

// 手机号注册
func MobileReg(c *gin.Context) {
	data := struct {
		Mobile string `json:"mobile"`
	}{}
	if err := c.BindJSON(&data); err != nil {
		e.ApiErr(c, err.Error())
		return
	}

	if _, err := models.UsersCheckMobile(data.Mobile); err == nil {
		e.ApiErr(c, "已经注册，请勿重复注册")
		return
	} else {
		//产生注册
		model := models.UcenterUsers{
			Mobile:   data.Mobile,
			Nickname: "注册用户",
			Avatar:   "https://image.ddgongjiang.com/FrIcNLhH4VlK67Nhu_m-GtU-sbOh",
			Gender:   "保密",
			Status:   1,
			UserKey:  models.GetUserKey(),
		}
		models.UsersAdd(&model)

		user, _ := models.GetUsersInfoCuid(model.Id)

		e.ApiOk(c, "注册成功", struct {
			UserInfo models.UcenterUsers `json:"user_info"`
		}{
			UserInfo: *user,
		})
	}
}
