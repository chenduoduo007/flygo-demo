package web

import (
	"github.com/chenduoduo007/flygo"
	"github.com/chenduoduo007/flygo-demo/pkg/e"
)

type Flygo struct {
	C *flygo.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 封装 flygo.JSON
func (g *Flygo) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
