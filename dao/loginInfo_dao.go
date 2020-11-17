package dao

import (
	"github.com/go-xorm/xorm"
	"oldplayerback/models"
	"time"
)

type LoginInfoDao struct {
	engine *xorm.Engine
}

func NewLoginInfoDao(engine *xorm.Engine) *LoginInfoDao {
	return &LoginInfoDao{
		engine: engine,
	}
}

//CheckLoginTime 返回用户最后登录的时间
func (d *LoginInfoDao) CheckLoginTime(uid string, partnerId int) time.Time {
	data := &models.GUserLogininfo{
		Userid:    uid,
		Partnerid: partnerId,
	}

	ok, err := d.engine.
		Desc("Logintime").
		Limit(1).
		Get(data)

	if ok && err == nil {
		return data.Logintime
	} else {
		return time.Time{}
	}
}
