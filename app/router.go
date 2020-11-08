package app

import (
	"github.com/chenduoduo007/flygo"
	"github.com/chenduoduo007/flygo-demo/app/api/v1"
	"github.com/chenduoduo007/flygo-demo/minddleware/errorhandler"
	"github.com/chenduoduo007/flygo-demo/pkg/e"
)

func InitRouter() *flygo.Engine  {
	r := flygo.New()
	r.Use(flygo.Logger())
	r.Use(flygo.Recovery())
	apiv1 := r.Group("/api/v1")
	//apiv1.Use(func_run_time.FuncRunTime())
	apiv1.Use(errorhandler.FrameworkError())
	{
		//获取标签列表
		apiv1.POST("/app_test", v1.ApiTest)
	}
	r.NoRoute(func(c *flygo.Context) {
		c.JSON(e.NOT_FOUND, flygo.H{
			"code": e.NOT_FOUND,
			"msg":  e.GetMsg(e.NOT_FOUND),
			"data": map[string]interface{}{},
		})
	})
	return r
}