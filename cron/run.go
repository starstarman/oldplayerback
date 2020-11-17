package cron

import (
	"log"
	"oldplayerback/web/controllers"
	"time"
)

func ClearMap() {
	//每隔1天清除下缓存
	time.AfterFunc(time.Minute*2, ClearMap)
	controllers.Sm.Range(func(key interface{}, value interface{}) bool {
		controllers.Sm.Delete(key)
		return true
	})
	log.Println("清理缓存")
}
