package routers

import (
	"newsWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandleReg")
    beego.Router("/login",&controllers.UserController{},"get:Showlogin;post:HandleLogin")
}
