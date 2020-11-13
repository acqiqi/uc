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

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数
	log.Println("Init Project")
	setting.Setup() //处理配置文件
	models.Setup()  //模型
	logging.Setup() //日志
	gredis.Setup()  // redis
	//utils.InitModel()
	log.Println("Ojbk")
	//WorkIn()

	list, err := models.MbGetAllList()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(list)

	works, err := models.UsersWorkGetSyncList()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(works)
	for i, v := range list {
		models.UsersWorkEdit(works[i].Id, map[string]interface{}{
			"avatar": v.FaceImageUrl,
		})
	}

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
