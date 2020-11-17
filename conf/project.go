package conf

import "time"

const SysTimeform = "2006-01-02 15:04:05"

// 中国时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

var Day = 86400

//设置玩家回归天数 30天
var OldPlayerBackDay = 30
