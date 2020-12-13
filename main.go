package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"uc/lib/gredis"
	"uc/lib/logging"
	"uc/lib/setting"
	"uc/models"
	"uc/router"
)

type Link struct {
	Content   string `json:"content"`
	Unit      string `json:"unit"`
	Zl        string `json:"zl"`
	Fl        string `json:"fl"`
	Num       int    `json:"num"`
	Work      string `json:"work"`
	UnitTotal string `json:"unit_total"`
	Total     int    `json:"total"`
	Describe  string `json:"describe"`
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数
	log.Println("Init Project")
	setting.Setup() //处理配置文件
	models.Setup()  //模型
	logging.Setup() //日志
	gredis.Setup()  // redis
	//utils.InitModel()
	//WorkIn()
	list, err := models.OrdersLogGetGaiList()
	if err != nil {
		log.Println("made zhizhang")
	}
	for _, v := range list {
		//查询订单是否存在
		order, err := models.SmUsersServiceGetInfoByOrderId(v.OrderId)
		if err == nil {
			log.Println(order.Id)
			models.OrdersLogEdit(v.Id, map[string]interface{}{
				"project_id": order.Id,
			})
		}
	}

	//log.Println(list)
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := router.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
