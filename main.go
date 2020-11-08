package main

import (
	"fmt"
	"github.com/chenduoduo007/flygo-demo/app"
	"github.com/chenduoduo007/flygo-demo/pkg/logging"
	"github.com/chenduoduo007/flygo-demo/pkg/setting"
	"log"
	"net/http"
)

func  init()  {
	setting.Setup()
	logging.Setup()
}

func main()  {
	routersInit := app.InitRouter()
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
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("[error] %s", err)
	}


	// 设置运行时最大核数
	//num := runtime.NumCPU()
	//fmt.Println(num)
	//runtime.GOMAXPROCS(num)

	//r := flygo.Default()
	//r.GET("/", func(c *flygo.Context) {
	//	c.String(http.StatusOK, "Hello Flygo")
	//})
	//
	//r.GET("/api", func(c *flygo.Context) {
	//	data := map[string]interface{}{
	//		"code": 200,
	//		"msg": "success",
	//		"data": map[string]interface{},
	//	}
	//	c.JSON(http.StatusOK, data)
	//})
	//r.Run(":9999")

}

