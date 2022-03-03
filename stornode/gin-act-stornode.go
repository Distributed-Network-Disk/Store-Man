package stornode

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	stornodeip := c.Query("stornodeip")
	stornodelistenport := c.Query("stornodelistenport")
	AliveNode[stornodeip+":"+stornodelistenport] = time.Now()
	c.Writer.WriteString("pong")
	c.Done()
}
