package services

import (
	"oldplayerback/dao"
	"oldplayerback/datasource"
)

type UserBackService interface {
	SetBack(uid string, partnerId int) error
}

type userBackService struct {
	dao *dao.UserBackDao
}

func NewUserBackService() UserBackService {
	return &userBackService{
		dao: dao.NewUserBackDao(datasource.InstanceDbMaster()),
	}
}

func (i *userBackService) SetBack(uid string, partnerId int) error {
	return i.dao.SetBack(uid, partnerId)
}
