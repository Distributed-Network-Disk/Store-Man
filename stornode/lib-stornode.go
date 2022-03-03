package stornode

import (
	"Store-Man/config"
	"time"
)

var (
	AliveNode map[string]time.Time
	Config    config.CONF
)

func CheckAliveNode() {
	for k, v := range AliveNode {
		if time.Since(v).Seconds() > float64(Config.AliveTimeout) {
			delete(AliveNode, k)
		}
	}
}
