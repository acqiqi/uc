package auth

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"uc/lib/e"
	"uc/lib/gredis"
	"uc/lib/utils"
	"uc/models"
)

//手机自动登录
func MobileAutoReg(c *gin.Context) {
	data := struct {
		Mobile   string `json:"mobile"`
		Code     string `json:"code"`
		IsRe     int    `json:"is_re"`
		Password string `json:"password"`
	}{}
	if err := c.BindJSON(&data); err != nil {
		e.ApiErr(c, err.Error())
		return
	}

	if !utils.CheckMobile(data.Mobile) {
		e.ApiErr(c, "请输入正确的手机号")
		return
	}

	//检测验证码是否存在
	vali_code := gredis.GetCacheString("AUTOREG" + data.Mobile)
	if data.Code != vali_code || data.Code == "" {
		e.ApiErr(c, "请输入正确的验证码")
		return
	}

	//查询用户是否存在
	user, err := models.UsersCheckMobile(data.Mobile)
	if err == nil {
		// 已经注册
		//生成Token
		token, err := utils.GenerateToken(user.Id)
		if err != nil {
			e.ApiErr(c, "生成token失败，请联系管理员")
			return
		}

		if data.IsRe == 1 {
			re_data := make(map[string]interface{})
			re_data["password"] = utils.PasswordEncode(data.Password)
			models.UsersEdit(user.Id, re_data)
		}

		e.ApiOk(c, "登录成功", struct {
			Token string `json:"token"`
		}{Token: token})
	} else {
		// 未注册
		user_data := models.UcenterUsers{
			Mobile:       data.Mobile,
			Nickname:     "注册用户",
			Avatar:       "https://image.ddgongjiang.com/FrIcNLhH4VlK67Nhu_m-GtU-sbOh",
			Gender:       "保密",
			Status:       1,
			UserKey:      models.GetUserKey(),
			BindUserinfo: 1,
		}

		if data.IsRe == 1 {
			user_data.Password = utils.PasswordEncode(data.Password)
		}

		if err := models.UsersAdd(&user_data); err != nil {
			e.ApiErr(c, "创建用户失败")
			return
		}

		user, _ := models.GetUsersInfoCuid(user_data.Id)
		//生成Token
		token, err := utils.GenerateToken(user.Id)
		if err != nil {
			e.ApiErr(c, "生成token失败，请联系管理员")
			return
		}

		e.ApiOk(c, "登录成功", struct {
			Token string `json:"token"`
		}{Token: token})
	}

}

func MobilePasswordLogin(c *gin.Context) {
	data := struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}{}

	if err := c.BindJSON(&data); err != nil {
		e.ApiErr(c, err.Error())
		return
	}

	if !utils.CheckMobile(data.Mobile) {
		e.ApiErr(c, "请输入正确的手机号")
		return
	}

	//查询用户是否存在
	user, err := models.UsersCheckMobile(data.Mobile)
	if err == nil {

		//判断密码
		if !utils.PasswordValidate(data.Password, user.Password) {
			e.ApiErr(c, "密码不正确")
			return
		}

		//生成Token
		token, err := utils.GenerateToken(user.Id)
		if err != nil {
			e.ApiErr(c, "生成token失败，请联系管理员")
			return
		}
		e.ApiOk(c, "登录成功", struct {
			Token string `json:"token"`
		}{Token: token})
	} else {
		e.ApiErr(c, "账号或密码不正确")
	}
}

// 获取自动注册手机号验证码
func MobileGetRegSMS(c *gin.Context) {
	data := struct {
		Mobile string `json:"mobile"`
	}{}
	if err := c.BindJSON(&data); err != nil {
		e.ApiErr(c, err.Error())
		return
	}

	if !utils.CheckMobile(data.Mobile) {
		e.ApiErr(c, "请输入正确的手机号")
		return
	}
	// 生成验证码
	code := utils.RandInt64(100000, 999999)
	code_str := strconv.FormatInt(code, 10)
	gredis.SetCacheString("AUTOREG"+data.Mobile, code_str, 7200)
	utils.SendSMSLuosinao(data.Mobile, "你的验证码是："+code_str)
	e.ApiOk(c, "获取成功", e.GetEmptyStruct())
}

// 获取修改密码的短信
func MobileGetRPSMS(c *gin.Context) {
	data := struct {
		Mobile string `json:"mobile"`
	}{}
	if err := c.BindJSON(&data); err != nil {
		e.ApiErr(c, err.Error())
		return
	}

	if !utils.CheckMobile(data.Mobile) {
		e.ApiErr(c, "请输入正确的手机号")
		return
	}

	_, err := models.UsersCheckMobile(data.Mobile)
	if err != nil {
		e.ApiErr(c, "手机号不正确")
		return
	}

	// 生成验证码
	code := utils.RandInt64(100000, 999999)
	code_str := strconv.FormatInt(code, 10)
	gredis.SetCacheString("RP"+data.Mobile, code_str, 7200)
	utils.SendSMSLuosinao(data.Mobile, "你的验证码是："+code_str)
	e.ApiOk(c, "获取成功", e.GetEmptyStruct())
}

// 根据手机号修改密码
func MobileRePassword(c *gin.Context) {
	data := struct {
		Mobile     string `json:"mobile"`
		Code       string `json:"code"`
		Password   string `json:"password"`
		RePassword string `json:"repassword"`
	}{}
	if err := c.BindJSON(&data); err != nil {
		e.ApiErr(c, err.Error())
		return
	}
	if !utils.CheckMobile(data.Mobile) {
		e.ApiErr(c, "请输入正确的手机号")
		return
	}
	//检测验证码是否存在
	vali_code := gredis.GetCacheString("RP" + data.Mobile)
	if data.Code != vali_code || data.Code == "" {
		e.ApiErr(c, "请输入正确的验证码")
		return
	}

	if data.Password != data.RePassword {
		e.ApiErr(c, "重复密码有误")
		return
	}
	if len(data.Password) < 6 {
		e.ApiErr(c, "请输入至少6位密码")
		return
	}

	//查询用户是否存在
	user, err := models.UsersCheckMobile(data.Mobile)
	if err != nil {
		e.ApiErr(c, "手机号不正确")
		return
	}

	save_data := make(map[string]interface{})
	save_data["password"] = utils.PasswordEncode(data.Password)
	models.UsersEdit(user.Id, save_data)
	e.ApiOk(c, "修改密码成功", e.GetEmptyStruct())
}
