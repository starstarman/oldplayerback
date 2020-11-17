package comm

import (
	"oldplayerback/conf"
	"time"
)

// 当前时间的时间戳
func NowUnix() int {
	return int(time.Now().In(conf.SysTimeLocation).Unix())
}
