package v1

import (
	"github.com/chenduoduo007/flygo"
	"github.com/chenduoduo007/flygo-demo/pkg/e"
	"github.com/chenduoduo007/flygo-demo/pkg/web"
	"net/http"
)


func ApiTest(c *flygo.Context)  {
	appG := web.Flygo{C: c}
	// 模拟服务错误
	//goodsList := make([]int,0)
	//println(goodsList[2])

	res := map[string]interface{}{
		"goods_count": 100,
		"goods_list": []int{1,2,2,2,4},
	}
	appG.Response(http.StatusOK, e.SUCCESS, res)
}