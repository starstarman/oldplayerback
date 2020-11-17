package models

import (
	"time"
)

type GUserLogininfo struct {
	Userid    string    `xorm:"not null comment('用户Id') VARCHAR(20)"`
	Partnerid int       `xorm:"comment('合作商Id') INT(11)"`
	Playerid  string    `xorm:"comment('玩家Id') CHAR(36)"`
	Logintime time.Time `xorm:"comment('登录时间') DATETIME"`
}
