package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"oldplayerback/comm"
	"oldplayerback/conf"
	"oldplayerback/services"
	"sync"
)

//缓存
var Sm sync.Map
var Mu sync.Mutex

type User struct {
	Uid       string `json:"UserId"`
	PartnerId int    `json:"PartnerId"`
}

type ActivityController struct {
	Ctx              iris.Context
	ServiceLogininfo services.LoginInfoService
	ServiceUserback  services.UserBackService
}

//Back 是否能够参与老玩家回归活动
func (c *ActivityController) PostBack() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""

	uid := c.Ctx.PostValue("UserId")
	if uid == "" {
		rs["code"] = 4000
		rs["msg"] = "UserId param err"
		return rs
	}

	partnerId, err := c.Ctx.PostValueInt("PartnerId")
	if err != nil {
		rs["code"] = 4000
		rs["msg"] = "PartnerId param err:" + err.Error()
		return rs
	}

	//查询缓存如果回归过了 不去命中数据库直接返回
	if val, ok := Sm.Load(uid); ok {
		fmt.Println("命中缓存")
		if val == "back" {
			rs["code"] = 4002
			rs["msg"] = "您已回归请勿重复操作！"
			return rs
		}
		if val == "noback" {
			rs["code"] = 4001
			rs["msg"] = "不满足回归玩家条件"
			return rs
		}
		if val == "okback" {
			rs["code"] = 200
			rs["msg"] = "满足回归玩家条件!"
			return rs
		}
	}

	Mu.Lock()
	defer Mu.Unlock()
	//双检查，保证只命中一次数据库
	if val, ok := Sm.Load(uid); ok {
		fmt.Println("命中缓存")
		if val == "back" {
			rs["code"] = 4002
			rs["msg"] = "您已回归请勿重复操作！"
			return rs
		}
		if val == "noback" {
			rs["code"] = 4001
			rs["msg"] = "不满足回归玩家条件"
			return rs
		}
		if val == "okback" {
			rs["code"] = 200
			rs["msg"] = "满足回归玩家条件!"
			return rs
		}
	}

	time := c.ServiceLogininfo.CheckLoginTime(uid, partnerId)
	if (comm.NowUnix()-int(time.Unix()))/conf.Day < conf.OldPlayerBackDay {
		rs["code"] = 4001
		rs["msg"] = "不满足回归玩家条件"
		Sm.Store(uid, "noback")
		return rs
	} else {
		rs["code"] = 200
		rs["msg"] = "满足回归玩家条件!"
		Sm.Store(uid, "okback")
		return rs
	}
}

//Setback 设置回归
func (c *ActivityController) PostSetback() map[string]interface{} {
	rs := make(map[string]interface{})
	uid := c.Ctx.PostValue("UserId")
	if uid == "" {
		rs["code"] = 4000
		rs["msg"] = "UserId param err"
		return rs
	}

	partnerId, err := c.Ctx.PostValueInt("PartnerId")
	if err != nil {
		rs["code"] = 4000
		rs["msg"] = "PartnerId param err:" + err.Error()
		return rs
	}

	//检查是否满足回归条件
	c.PostBack()
	//在缓存中查询玩家是否已经回归，如果回归直接返回，不去命中数据库
	if val, ok := Sm.Load(uid); ok {
		fmt.Println("命中缓存")
		if val == "noback" {
			rs["code"] = 4001
			rs["msg"] = "不满足回归玩家条件"
			return rs
		}
		if val == "back" {
			rs["code"] = 4002
			rs["msg"] = "您已回归请勿重复操作！"
			return rs
		}
	}

	Mu.Lock()
	defer Mu.Unlock()

	err = c.ServiceUserback.SetBack(uid, partnerId)
	if err != nil {
		rs["code"] = 4002
		rs["msg"] = "您已回归请勿重复操作！"
		Sm.Store(uid, "back")
		return rs
	}

	rs["code"] = 200
	rs["msg"] = "回归成功！"
	//设置缓存
	Sm.Store(uid, "back")
	return rs
}
