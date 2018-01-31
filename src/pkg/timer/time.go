package timer

import (
	"Repo-watcher/src/pkg/config"
	"fmt"
	"time"
)

func GetCurrentJapanHourMin() string {
	// Time setting
	now := time.Now()
	nowUTC := now.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	nowJST := nowUTC.In(jst)
	hour, min, _ := nowJST.Clock()
	var hourMin string
	if min < 10 {
		hourMin = fmt.Sprintf("%d:0%d", hour, min)
	} else {
		hourMin = fmt.Sprintf("%d:%d", hour, min)
	}
	return hourMin
}

func IsTime(c *config.ConfigFile, t string) bool {
	for _, time := range c.AlarmTime {
		if time == t {
			return true
		}
	}
	return false
}
