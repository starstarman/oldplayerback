package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"oldplayerback/bootstrap"
	"oldplayerback/services"
	"oldplayerback/web/controllers"
)

//Configure register routes to the app
func Configure(b *bootstrap.Bootstrapper) {
	logininfoService := services.NewLoginInfoService()
	userbackService := services.NewUserBackService()

	oldplayer := mvc.New(b.Party("/player"))
	oldplayer.Register(logininfoService, userbackService)
	oldplayer.Handle(new(controllers.ActivityController))
}
