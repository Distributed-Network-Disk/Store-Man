package main

import (
	"Store-Man/client"
	"Store-Man/config"
	"Store-Man/stornode"
	"flag"
	"fmt"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	Config config.CONF
)

func main() {
	var (
		Show_Banner bool
	)
	runtime.GOMAXPROCS(20)

	flag.IntVar(&Config.ListenPort, "port", 9999, "port")
	flag.IntVar(&Config.AliveTimeout, "timeout", 300, "timeout (s)")
	flag.BoolVar(&Show_Banner, "v", false, "show banner")
	flag.Parse()

	if Show_Banner {
		fmt.Println("Store Manager, Powered by Totoro and Golang")
		return
	}

	stornode.Config = Config

	gin.SetMode(gin.ReleaseMode)
	storman := gin.New()

	storman.GET("/ping", client.Ping)

	storman.GET("/nodeping", stornode.Ping)

	storman.Run(":" + strconv.Itoa(Config.ListenPort))
}
