package func_run_time

import (
	"github.com/chenduoduo007/flygo"
	"log"
	"time"
)

func FuncRunTime() flygo.HandlerFunc {
	return func(c *flygo.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		log.Print(c.HandlerName(),latency)
	}
}