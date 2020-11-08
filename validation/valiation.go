package validation

import (
	"github.com/chenduoduo007/flygo"
	"github.com/gin-gonic/gin/binding"
)

func FormValidation(c *flygo.Context,b interface{}){
	if err := c.ShouldBindWith(b, binding.Form); err != nil {
		mapErr := map[string]interface{}{
			"code": 400,
			"msg": err.Error(),
		}
		// 抛出异常，中断请求
		panic(mapErr)
	}
}

