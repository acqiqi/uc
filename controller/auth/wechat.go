package auth

import (
	"github.com/gin-gonic/gin"
	"uc/lib/e"
	"uc/lib/utils"
	"uc/lib/utils/go_wechat"
	"uc/models"
)

func GetMPOpenId(c *gin.Context) {
	data := struct {
		Code string `json:"code"`
		//PlatformKey string `json:"platform_key"`
	}{}
	if err := c.BindJSON(&data); err != nil {
		e.ApiErr(c, err.Error())
		return
	}

	platformKey := c.GetHeader("PlatformKey")
	platform, err := models.PlatformGetInfoOrKey(platformKey)
	if err != nil {
		e.ApiErr(c, "PlatformKey有误")
		return
	}

	cb, err := go_wechat.GetWeixinToken(data.Code, platform.Ak, platform.Sk)
	if err != nil {
		e.ApiErr(c, err.Error())
		return
	}

	e.ApiOk(c, "获取成功", cb)
}

// 小程序openid 快速登录
func SmallQuckLogin(c *gin.Context) {
	data := struct {
		Code string `json:"code"`
		//PlatformKey string `json:"platform_key"`
	}{}
	if err := c.BindJSON(&data); err != nil {
		e.ApiErr(c, err.Error())
		return
	}

	platformKey := c.GetHeader("PlatformKey")
	platform, err := models.PlatformGetInfoOrKey(platformKey)
	if err != nil {
		e.ApiErr(c, "PlatformKey有误")
		return
	}

	cb, err := go_wechat.SmallLogin(data.Code, platform.Ak, platform.Sk)
	if err != nil {
		e.ApiErr(c, err.Error())
		return
	}

	//查询是否注册
	user, err := models.UcenterOpenidCheckInfo(platformKey, cb.Openid)
	if err != nil {
		e.ApiOpt(c, e.API_NOT_AUTH_REG_CODE, e.GetMsg(e.API_NOT_AUTH_REG_CODE), e.GetEmptyStruct())
		return
	}

	//生成Token
	token, err := utils.GenerateToken(user.User.Id)
	if err != nil {
		e.ApiErr(c, "生成token失败，请联系管理员")
		return
	}

	e.ApiOk(c, "登录成功", struct {
		Token string `json:"token"`
	}{Token: token})
}

// 使用微信小程序的获取手机号方式来注册
func SmallQuckMobileReg(c *gin.Context) {
	data := struct {
		Code          string `json:"code"`
		EncryptedData string `json:"encrypted_data"`
		Iv            string `json:"iv"`
	}{}
	if err := c.BindJSON(&data); err != nil {
		e.ApiErr(c, err.Error())
		return
	}
	platformKey := c.GetHeader("PlatformKey")
	platform, err := models.PlatformGetInfoOrKey(platformKey)
	if err != nil {
		e.ApiErr(c, "PlatformKey有误")
		return
	}
	// 获取openid
	cb, err := go_wechat.SmallLogin(data.Code, platform.Ak, platform.Sk)
	if err != nil {
		e.ApiErr(c, err.Error())
		return
	}

	wx := go_wechat.NewWXUserDataCrypt(data.EncryptedData, cb.SessionKey)
	cb_mobile, err := wx.DecryptUserMobile(data.EncryptedData, data.Iv)
	if err != nil {
		e.ApiErr(c, err.Error())
		return
	}

	if !utils.CheckMobile(cb_mobile.PurePhoneNumber) {
		e.ApiErr(c, "手机号不正确")
		return
	}

	//判断是否注册
	user, err := models.UsersCheckMobile(cb_mobile.PurePhoneNumber)
	if err == nil {
		// 用户注册
		// 更新openid
		if !models.UcenterOpenidUp(user.Id, cb.Openid, platformKey, "wechat") {
			e.ApiErr(c, "更新Openid失败")
			return
		}
		//生成Token
		token, err := utils.GenerateToken(user.Id)
		if err != nil {
			e.ApiErr(c, "生成token失败，请联系管理员")
			return
		}
		e.ApiOk(c, "登录成功", struct {
			Token    string              `json:"token"`
			UserInfo models.UcenterUsers `json:"user_info"`
		}{
			Token:    token,
			UserInfo: *user,
		})
	} else {
		// 用户未注册
		data := models.UcenterUsers{
			Username: "",
			Password: "",
			Mobile:   cb_mobile.PurePhoneNumber,
			Nickname: "注册用户",
			Email:    "",
			Avatar:   "https://image.ddgongjiang.com/FrIcNLhH4VlK67Nhu_m-GtU-sbOh",
			Gender:   "保密",
			UserKey:  models.GetUserKey(),
		}
		if err := models.UsersAdd(&data); err != nil {
			e.ApiErr(c, "创建用户失败")
			return
		}
		if u, err := models.GetUsersInfoCuid(data.Id); err == nil {
			//生成Token
			token, err := utils.GenerateToken(u.Id)
			if err != nil {
				e.ApiErr(c, "生成token失败，请联系管理员")
				return
			}
			e.ApiOk(c, "登录成功", struct {
				Token    string              `json:"token"`
				UserInfo models.UcenterUsers `json:"user_info"`
			}{
				Token:    token,
				UserInfo: *user,
			})
		} else {
			e.ApiErr(c, "创建用户失败!")
		}
	}
}

// 绑定用户信息
func SmallBindUserInfoLogin(c *gin.Context) {
	data := struct {
		Avatar   string `json:"avatar"`
		Nickname string `json:"nickname"`
		Code     string `json:"code"`
	}{}
	if err := c.BindJSON(&data); err != nil {
		e.ApiErr(c, err.Error())
		return
	}

	if data.Avatar == "" || data.Nickname == "" {
		e.ApiErr(c, "参数不正确")
		return
	}

	platformKey := c.GetHeader("PlatformKey")
	platform, err := models.PlatformGetInfoOrKey(platformKey)
	if err != nil {
		e.ApiErr(c, "PlatformKey有误")
		return
	}
	// 获取openid
	cb, err := go_wechat.SmallLogin(data.Code, platform.Ak, platform.Sk)
	if err != nil {
		e.ApiErr(c, err.Error())
		return
	}

	// 检测openid 是否注册
	user, err := models.UcenterOpenidCheckInfo(platformKey, cb.Openid)
	if err != nil {
		e.ApiOpt(c, e.API_NOT_AUTH_REG_CODE, e.GetMsg(e.API_NOT_AUTH_REG_CODE), e.GetEmptyStruct())
		return
	}

	user.User.Avatar = data.Avatar
	user.User.Nickname = data.Nickname
	user.User.BindUserinfo = 1
	if err := models.UsersEdit(user.User.Id, user.User); err != nil {
		e.ApiErr(c, "编辑失败")
		return
	}
	//生成Token
	token, err := utils.GenerateToken(user.User.Id)
	if err != nil {
		e.ApiErr(c, "生成token失败，请联系管理员")
		return
	}
	e.ApiOk(c, "登录成功", struct {
		Token    string              `json:"token"`
		UserInfo models.UcenterUsers `json:"user_info"`
	}{
		Token:    token,
		UserInfo: user.User,
	})
}
