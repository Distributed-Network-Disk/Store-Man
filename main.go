package main

// Copyright [2022] [totoroyyw]

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
