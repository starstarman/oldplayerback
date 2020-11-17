package bootstrap

import (
	"github.com/kataras/iris/v12"
	"log"
	"oldplayerback/cron"
	"time"
)

type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
}

type Configurator func(*Bootstrapper)

func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}

// 启动计划任务服务
func (b *Bootstrapper) setupCron() {
	//清楚缓存任务
	go cron.ClearMap()
	log.Println("开启缓存任务")
}

func (b *Bootstrapper) Bootstrap() {
	b.setupCron()
}
