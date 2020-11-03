package main

import (
	"log"
	"strconv"
	"time"
	"uc/lib/utils"
	"uc/models"
)

func WorkIn() {
	list, err := models.MaGetAllList()
	if err != nil {
		log.Println(err.Error())
		return
	}

	for i, v := range list {
		if !utils.InArray(models.WorkType, v.WorkerType) {
			log.Println("NOJbk", i, v)
			continue
		}
		if v.IdNo == "" {
			log.Println("没有身份证")
			continue
		}
		if v.Mobile == "" {
			log.Println("NO Mobile")
			continue
		}
		_, err := models.UsersCheckMobile(v.Mobile)
		if err == nil {
			log.Println("注册过了")
			continue
		}
		uc_model := models.UcenterUsers{
			Mobile:   v.Mobile,
			Nickname: v.Nickname,
			Email:    "",
			Avatar:   "https://image.ddgongjiang.com/FrIcNLhH4VlK67Nhu_m-GtU-sbOh",
			Gender:   "保密",
			Status:   1,
			UserKey:  models.GetUserKey(),
		}
		models.UsersAdd(&uc_model)

		work_model := models.UcenterUsersWork{
			Cuid:              uc_model.Id,
			Name:              v.Nickname,
			CardId:            v.IdNo,
			CardTop:           v.IdCardzmurl,
			CardRen:           v.IdCardfmurl,
			Address:           "",
			Nation:            "汉族",
			Gender:            "男",
			Province:          "广东省",
			City:              "深圳市",
			Area:              "",
			IsAuth:            0,
			WorkType:          v.WorkerType,
			LocationAddress:   "深圳",
			LocationProvince:  "广东省",
			LocationCity:      "深圳市",
			LocationArea:      "福田区",
			LocationLongitude: 0,
			LocationLatitude:  0,
			OfficeMobile:      v.Mobile,
			Miaoshu:           models.DescArray[utils.RandInt64(1, 115)],
			GongzuoTime:       "",
			IsGongren:         1,
			Birth:             utils.Time{BuildSR(v.IdNo)},
			IsGongzuo:         0,
			GongrenType:       "",
			Home:              "",
		}
		models.UsersWorkAdd(&work_model)
	}

	log.Println("哦几把剋")

}

func BuildSR(id_card string) time.Time {
	idCard := id_card
	idcard_year, _ := strconv.Atoi(utils.Substr(idCard, 6, 4)) // 年
	idcard_mo, _ := strconv.Atoi(utils.Substr(idCard, 10, 2))  // 月
	idcard_day, _ := strconv.Atoi(utils.Substr(idCard, 12, 2)) // 日
	test_t := time.Time{}
	cb_time := test_t.AddDate(idcard_year-1, idcard_mo-1, idcard_day-1)
	log.Println(cb_time)
	return cb_time
}
