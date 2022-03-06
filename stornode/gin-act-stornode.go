package stornode

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
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	stornodeip := c.Query("stornodeip")
	stornodelistenport := c.Query("stornodelistenport")
	AliveNode[stornodeip+":"+stornodelistenport] = time.Now()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "pong"})
}
