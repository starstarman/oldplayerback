package services

import (
	"oldplayerback/dao"
	"oldplayerback/datasource"
	"time"
)

type LoginInfoService interface {
	CheckLoginTime(uid string, partnerId int) time.Time
}

type loginInfoService struct {
	dao *dao.LoginInfoDao
}

func NewLoginInfoService() LoginInfoService {
	return &loginInfoService{
		dao: dao.NewLoginInfoDao(datasource.InstanceDbMaster()),
	}
}

func (i *loginInfoService) CheckLoginTime(uid string, partnerId int) time.Time {
	return i.dao.CheckLoginTime(uid, partnerId)
}
