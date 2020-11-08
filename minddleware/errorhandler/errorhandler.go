package errorhandler

import (
	"github.com/chenduoduo007/flygo"
	"github.com/chenduoduo007/flygo-demo/pkg/e"
	"github.com/chenduoduo007/flygo-demo/pkg/logging"
	"github.com/chenduoduo007/flygo-demo/pkg/setting"
	"reflect"
)

func FrameworkError() flygo.HandlerFunc {
	return func(c *flygo.Context) {
		defer func() {
			if r := recover(); r != nil {
				// 处理自定义错误
				if reflect.TypeOf(r).Kind() == reflect.Map {
					mapRec := r.(map[string]interface{})
					if mapRec["code"] == e.PARAMS_ERROR {
						c.JSON(e.PARAMS_ERROR, flygo.H{
							"code": e.PARAMS_ERROR,
							"msg":  mapRec["msg"],
							"data": map[string]interface{}{},
						})
						//c.Abort()
						return
					}
				}
				if setting.ServerSetting.RunMode == "release" {
					//	线上记录到log中,记录完，返回服务错误]
					logging.Error(r)
					c.JSON(e.ERROR, flygo.H{
						"code": e.ERROR,
						"msg":  e.GetMsg(e.ERROR),
						"data": map[string]interface{}{},
					})
					//c.Abort()
				} else {
					c.JSON(e.ERROR, flygo.H{
						"code": e.ERROR,
						"msg": errorToString(r),
						"data": map[string]interface{}{},
					})
					//c.Abort()
					logging.Error(r)
					panic(r)
				}
			}
		}()
		c.Next()
	}
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}