package dao

import (
	"github.com/go-xorm/xorm"
	"oldplayerback/models"
)

type UserBackDao struct {
	engine *xorm.Engine
}

func NewUserBackDao(engine *xorm.Engine) *UserBackDao {
	return &UserBackDao{
		engine: engine,
	}
}

//SetBack 设置用户回归
func (d *UserBackDao) SetBack(uid string, partnerId int) error {
	data := &models.GUserBack{
		Userid:    uid,
		Partnerid: partnerId,
	}
	_, err := d.engine.Insert(data)
	return err
}
