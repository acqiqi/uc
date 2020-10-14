package v1

import (
	"github.com/gin-gonic/gin"
	"uc/lib/e"
	"uc/models"
)

// userinfo
func GetUserInfo(c *gin.Context) {
	user_info, _ := c.Get("user_info")
	user, err := models.GetUsersInfoCuid(user_info.(models.UcenterUsers).Id)
	if err != nil {
		e.ApiErr(c, "非法请求")
		return
	}
	e.ApiOk(c, "获取成功", struct {
		UserInfo models.UcenterUsers `json:"user_info"`
	}{
		UserInfo: *user,
	})
}
